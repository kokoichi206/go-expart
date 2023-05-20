package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"syscall"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "ch1", "set run router mode")
	flag.Parse()

	runChap(mode)
}

// Global変数。
var netDeviceList []*netDevice

func runChap(mode string) {
	events := make([]syscall.EpollEvent, 10)

	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		log.Fatalf("failed to create epoll: %s", err)
	}

	interfaces, _ := net.Interfaces()
	for _, netif := range interfaces {
		fmt.Printf("netif.Name: %v\n", netif.Name)
		fmt.Printf("netif.Flags: %v\n", netif.Flags)

		if isIgnoreInterfaces(netif.Name) {
			continue
		}

		// どのレイヤー通信をしたいか
		// cf. TCP:
		//      AF_INET (IPv4 を意味する )
		//      SOCK_STREAM (TCP を意味する )
		sock, err := syscall.Socket(
			syscall.AF_PACKET,
			syscall.SOCK_RAW,
			int(htons(syscall.ETH_P_ALL)),
		)
		if err != nil {
			log.Fatalf("failed to create socket: %s", err)
		}

		// NW インタフェースと紐付け
		addr := syscall.SockaddrLinklayer{
			Protocol: htons(syscall.ETH_P_ALL),
			Ifindex:  netif.Index,
		}
		if err := syscall.Bind(sock, &addr); err != nil {
			log.Fatalf("failed to bind: %s", err)
		}

		fmt.Printf("created device %s sodket %d address %s\n", netif.Name, sock, netif.HardwareAddr)

		// ソケットを作成するとデフォルトではブロッキングモードになるので、
		// パケットを受信するまで後続処理がぶろくされる。
		// ノンブロッキングに設定するのではなく、epoll の監視対象とする。
		if err := syscall.EpollCtl(
			epfd, syscall.EPOLL_CTL_ADD, sock,
			&syscall.EpollEvent{
				Events: syscall.EPOLLIN,
				Fd:     int32(sock),
			},
		); err != nil {
			log.Fatalf("afiled to epol ctl: %s", err)
		}

		// ノンブロッキングに設定
		// err = syscall.SetNonblock(sock, true)

		// これが必要
		// ないと『interrupted system call』が発生した
		syscall.SetNonblock(sock, false)

		netAddr, err := netif.Addrs()
		if err != nil {
			log.Fatalf("failed to get addr from NIC interface: %s", err)
		}

		netDev := &netDevice{
			name:       netif.Name,
			macaddr:    setMacAddr(netif.HardwareAddr),
			socket:     sock,
			socketaddr: addr,
			ipdev:      getIPdevice(netAddr),
		}

		netDeviceList = append(netDeviceList, netDev)
	}

	for {
		// パケットの受信を待つ！
		nfds, err := syscall.EpollWait(epfd, events, -1)
		if err != nil {
			// log.Fatalf("failed to epoll wait: %s", err)
			fmt.Printf("failed to epoll wait: %s\n", err)
		}

		for i := 0; i < nfds; i++ {
			for _, netdev := range netDeviceList {
				// イベントがあったソケットとマッチしたら、
				// パケットを読み込む処理を実行！
				if events[i].Fd == int32(netdev.socket) {
					if err := netdev.netDevicePoll(mode); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
}

const (
	connected ipRouteType = iota
	network
)

type ipRouteType uint8

type ipRouteEntry struct {
	iptype  ipRouteType
	netdev  *netDevice
	nexthop uint32
}

var IGNORE_INTERFACES = []string{"lo", "bond0", "dummy0", "tunl0", "sit0"}

func isIgnoreInterfaces(name string) bool {
	for _, v := range IGNORE_INTERFACES {
		if v == name {
			return true
		}
	}
	return false
}

// ETH_P_ALL などについて
// https://linuxjm.osdn.jp/html/LDP_man-pages/man7/packet.7.html
func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func setMacAddr(macAddrByte []byte) [6]uint8 {
	var macAddrUint8 [6]uint8
	for i, v := range macAddrByte {
		macAddrUint8[i] = v
	}

	return macAddrUint8
}

func byteToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

// ネットデバイスの受信処理
func (netdev *netDevice) netDevicePoll(mode string) error {
	recvbuffer := make([]byte, 1500)
	n, _, err := syscall.Recvfrom(netdev.socket, recvbuffer, 0)
	if err != nil {
		if n == -1 {
			return nil
		}

		return fmt.Errorf("recv err, n is %d, device is %s, err is %s", n, netdev.name, err)
	}

	if mode == "ch1" {
		fmt.Printf("Received %d bytes from %s: %x\n", n, netdev.name, recvbuffer[:n])
	} else {
		ethernetInput(netdev, recvbuffer[:n])
	}

	return nil
}

const ETHER_TYPE_IP uint16 = 0x0800
const ETHER_TYPE_ARP uint16 = 0x0806
const ETHER_TYPE_IPV6 uint16 = 0x86dd
const ETHERNET_ADDRES_LEN = 6

var ETHERNET_ADDRESS_BROADCAST = [6]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

func ethernetInput(netdev *netDevice, packet []byte) {
	// イーサネットのフレームとして解釈する！
	// 6-6-2 byte で Ethernet ヘッダがあり、
	// 宛先 Mac アドレス、送信元 Mac アドレス、タイプである。
	netdev.etheHeader.destAddr = setMacAddr(packet[:6])
	netdev.etheHeader.srcAddr = setMacAddr(packet[6:12])
	netdev.etheHeader.etherType = byteToUint16(packet[12:14])

	// 自分の Mac アドレス宛 or ブロードキャスト通信 でなければ処理を行わない。
	if (netdev.macaddr != netdev.etheHeader.destAddr) &&
		(netdev.etheHeader.destAddr != ETHERNET_ADDRESS_BROADCAST) {
		return
	}

	// イーサタイプの値から、上位プロトコルを特定する。
	switch netdev.etheHeader.etherType {
	case ETHER_TYPE_ARP:
		if err := arpInput(netdev, packet[14:]); err != nil {
			log.Println(err)
		}

	case ETHER_TYPE_IP:
		ipInput(netdev, packet[14:])
	}
}

const IP_ADDRESS_LEN = 4
const IP_ADDRESS_LIMITED_BROADCAST uint32 = 0xffffffff
const IP_PROTOCOL_NUM_ICMP uint8 = 0x01 // IP ヘッダの protocol(1 byte) が 01
const IP_PROTOCOL_NUM_TCP uint8 = 0x06  // IP ヘッダの protocol(1 byte) が 06
const IP_PROTOCOL_NUM_UDP uint8 = 0x11  // IP ヘッダの protocol(1 byte) が 11

type arpIPToEthernet struct {
	hardwareType        uint16 // Ethernet の場合は 1
	protocolType        uint16 // IPv4 の場合は 0x0800
	hardwareLen         uint8  // Mac アドレスの長さなら 6
	protocolLen         uint8  // IPv4 の長さなら 4
	opcode              uint16 // ARP Request 0c0001, ARP Reply 0c0002
	senderHardwareAddr  [6]uint8
	senderIPAddr        uint32
	targetHardwareAddrr [6]uint8
	targetIPAddr        uint32
}

func (arpmsg arpIPToEthernet) ToPacket() []byte {
	var b bytes.Buffer

	b.Write(uint16ToByte(arpmsg.hardwareType))
	b.Write(uint16ToByte(arpmsg.protocolType))
	b.Write([]byte{arpmsg.hardwareLen})
	b.Write([]byte{arpmsg.protocolLen})
	b.Write(uint16ToByte(arpmsg.opcode))
	b.Write(macToByte(arpmsg.senderHardwareAddr))
	b.Write(uint32ToByte(arpmsg.senderIPAddr))
	b.Write(macToByte(arpmsg.targetHardwareAddrr))
	b.Write(uint32ToByte(arpmsg.targetIPAddr))

	return b.Bytes()
}

func macToByte(macaddr [6]uint8) (b []byte) {
	for _, v := range macaddr {
		b = append(b, v)
	}
	return b
}

/*
ARPパケットの受信処理
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/arp.cpp#L139
*/
func arpInput(netdev *netDevice, packet []byte) error {
	// ARPパケットの規定より短かった場合。
	if len(packet) < 28 {
		fmt.Printf("received ARP Packet is too short")

		return fmt.Errorf("received ARP Packet is too short: size (%s)", len(packet))
	}

	arpMsg := arpIPToEthernet{
		hardwareType:        byteToUint16(packet[0:2]),
		protocolType:        byteToUint16(packet[2:4]),
		hardwareLen:         packet[4],
		protocolLen:         packet[5],
		opcode:              byteToUint16(packet[6:8]),
		senderHardwareAddr:  setMacAddr(packet[8:14]),
		senderIPAddr:        byteToUint32(packet[14:18]),
		targetHardwareAddrr: setMacAddr(packet[18:24]),
		targetIPAddr:        byteToUint32(packet[24:28]),
	}

	switch arpMsg.protocolType {
	case ETHER_TYPE_IP:

		if arpMsg.hardwareLen != ETHERNET_ADDRES_LEN {
			fmt.Println("Illegal hardware address length")

			return fmt.Errorf("Illegal hardware address length: %d", arpMsg.hardwareLen)
		}

		if arpMsg.protocolLen != IP_ADDRESS_LEN {
			fmt.Println("Illegal protocol address length")

			return fmt.Errorf("Illegal protocol address length: %d", arpMsg.protocolLen)
		}

		// オペレーションコードによって分岐！！
		if arpMsg.opcode == ARP_OPERATION_CODE_REQUEST {
			// ARPリクエストの受信。
			fmt.Printf("ARP Request Packet is %+v\n", arpMsg)
			arpRequestArrives(netdev, arpMsg)
		} else {
			// ARPリプライの受信（自身が ARP リクエストを投げたものの返信）。
			fmt.Printf("ARP Reply Packet is %+v\n", arpMsg)
			arpReplyArrives(netdev, arpMsg)
		}
	}

	return nil
}

func byteToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

const ARP_OPERATION_CODE_REQUEST = 1
const ARP_OPERATION_CODE_REPLY = 2
const ARP_HTYPE_ETHERNET uint16 = 0001

/*
ARPリクエストパケットの受信処理
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/arp.cpp#L181
*/
func arpRequestArrives(netdev *netDevice, arp arpIPToEthernet) {
	// IPアドレスが設定されているデバイスからの受信かつ要求されているアドレスが自分の物だったら
	fmt.Println("arpRequestArrives...")
	fmt.Printf("(netdev.ipdev.address != 00000000): %v\n", (netdev.ipdev.address != 00000000))
	fmt.Printf("(netdev.ipdev.address == arp.targetIPAddr): %v\n", (netdev.ipdev.address == arp.targetIPAddr))
	if netdev.ipdev.address != 00000000 && netdev.ipdev.address == arp.targetIPAddr {
		fmt.Printf("Sending arp reply to %s\n", printIPAddr(arp.targetIPAddr))
		// APRリプライのパケットを作成
		arpPacket := arpIPToEthernet{
			hardwareType:        ARP_HTYPE_ETHERNET,
			protocolType:        ETHER_TYPE_IP,
			hardwareLen:         ETHERNET_ADDRES_LEN,
			protocolLen:         IP_ADDRESS_LEN,
			opcode:              ARP_OPERATION_CODE_REPLY,
			senderHardwareAddr:  netdev.macaddr,
			senderIPAddr:        netdev.ipdev.address,
			targetHardwareAddrr: arp.senderHardwareAddr,
			targetIPAddr:        arp.senderIPAddr,
		}.ToPacket()

		// ethernetでカプセル化して送信
		ethernetOutput(netdev, arp.senderHardwareAddr, arpPacket, ETHER_TYPE_ARP)
	}
}

func printIPAddr(ip uint32) string {
	ipbyte := uint32ToByte(ip)
	return fmt.Sprintf("%d.%d.%d.%d", ipbyte[0], ipbyte[1], ipbyte[2], ipbyte[3])
}

// イーサネットにカプセル化して送信する。
func ethernetOutput(netdev *netDevice, destaddr [6]uint8, packet []byte, ethType uint16) {
	ethHeaderPacket := ethernetHeader{
		destAddr:  destaddr,
		srcAddr:   netdev.macaddr,
		etherType: ethType,
	}.ToPacket()

	// パケットのカプセル化！
	// イーサネットヘッダに、受け取った上位レイヤのパケットをつなげる。
	ethHeaderPacket = append(ethHeaderPacket, packet...)

	// ネットワークデバイスに送信する。
	if err := netdev.netDeviceTransmit(ethHeaderPacket); err != nil {
		log.Fatalf("netDeviceTransmit is err : %v", err)
	}
}

// NW インタフェースに bind したソケットに sendto でパケットが送信される。
func (netdev netDevice) netDeviceTransmit(data []byte) error {
	fmt.Println("netDeviceTransmiting...")
	err := syscall.Sendto(netdev.socket, data, 0, &netdev.socketaddr)
	if err != nil {
		return fmt.Errorf("failed to execute syscall sendTo: %w", err)
	}

	return nil
}

/*
ARPリプライパケットの受信処理
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/arp.cpp#L213
*/
func arpReplyArrives(netdev *netDevice, arp arpIPToEthernet) {
	// IPアドレスが設定されているデバイスからの受信だった場合。
	if netdev.ipdev.address != 00000000 {
		fmt.Printf("Added arp table entry by arp reply (%s => %s)\n", printIPAddr(arp.senderIPAddr), printMacAddr(arp.senderHardwareAddr))
		// ARPテーブルエントリに追加する。
		addArpTableEntry(netdev, arp.senderIPAddr, arp.senderHardwareAddr)
	}
}

func printMacAddr(macddr [6]uint8) string {
	var str string
	for _, v := range macddr {
		str += fmt.Sprintf("%x:", v)
	}
	return strings.TrimRight(str, ":")
}

type ipHeader struct {
	version        uint8
	headerLen      uint8
	tos            uint8  // Type of Service
	totalLen       uint16 // Total のパケット長
	identify       uint16 // 識別番号
	fragOffset     uint16 // フラグ
	ttl            uint8
	protocol       uint8  // 上位のプロトコル番号
	headerChecksum uint16 // ヘッダのチェックサム
	srcAddr        uint32
	destAddr       uint32
}

func (ipheader ipHeader) ToPacket(calc bool) (ipHeaderByte []byte) {
	var b bytes.Buffer

	b.Write([]byte{ipheader.version<<4 + ipheader.headerLen})
	b.Write([]byte{ipheader.tos})
	b.Write(uint16ToByte(ipheader.totalLen))
	b.Write(uint16ToByte(ipheader.identify))
	b.Write(uint16ToByte(ipheader.fragOffset))
	b.Write([]byte{ipheader.ttl})
	b.Write([]byte{ipheader.protocol})
	b.Write(uint16ToByte(ipheader.headerChecksum))
	b.Write(uint32ToByte(ipheader.srcAddr))
	b.Write(uint32ToByte(ipheader.destAddr))

	// checksumを計算する
	if calc {
		ipHeaderByte = b.Bytes()
		checksum := calcChecksum(ipHeaderByte)
		// checksumをセット
		ipHeaderByte[10] = checksum[0]
		ipHeaderByte[11] = checksum[1]
	} else {
		ipHeaderByte = b.Bytes()
	}

	return ipHeaderByte
}

func calcChecksum(packet []byte) []byte {
	// まず16ビット毎に足す
	sum := sumByteArr(packet)
	// あふれた桁を足す
	sum = (sum & 0xffff) + sum>>16
	// 論理否定を取った値をbyteにして返す
	return uint16ToByte(uint16(sum ^ 0xffff))
}

func sumByteArr(packet []byte) (sum uint) {
	for i, _ := range packet {
		if i%2 == 0 {
			sum += uint(byteToUint16(packet[i:]))
		}
	}
	return sum
}

func uint16ToByte(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func uint32ToByte(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

type arpTableEntry struct {
	macAddr [6]uint8
	ipAddr  uint32
	netdev  *netDevice
}

/**
 * ARPテーブル
 * グローバル変数にテーブルを保持
 */
var ArpTableEntryList []arpTableEntry

/*
ARPテーブルにエントリの追加と更新
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/arp.cpp#L23
*/
func addArpTableEntry(netdev *netDevice, ipaddr uint32, macaddr [6]uint8) {

	// 既存のARPテーブルの更新が必要か確認
	if len(ArpTableEntryList) != 0 {
		for _, arpTable := range ArpTableEntryList {
			// IPアドレスは同じだがMacアドレスが異なる場合は更新
			if arpTable.ipAddr == ipaddr && arpTable.macAddr != macaddr {
				arpTable.macAddr = macaddr
			}
			// Macアドレスは同じだがIPアドレスが変わった場合は更新
			if arpTable.macAddr == macaddr && arpTable.ipAddr != ipaddr {
				arpTable.ipAddr = ipaddr
			}
			// 既に存在する場合はreturnする
			if arpTable.macAddr == macaddr && arpTable.ipAddr == ipaddr {
				return
			}
		}
	}

	ArpTableEntryList = append(ArpTableEntryList, arpTableEntry{
		macAddr: macaddr,
		ipAddr:  ipaddr,
		netdev:  netdev,
	})
}

/*
IPパケットの受信処理
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/ip.cpp#L51
*/
func ipInput(inputdev *netDevice, packet []byte) {
	// IPアドレスのついていないインターフェースからの受信は無視。
	if inputdev.ipdev.address == 0 {
		return
	}

	// IPヘッダ長 = 20 bytes より短かったらドロップ。
	if len(packet) < 20 {
		fmt.Printf("Received IP packet too short from %s\n", inputdev.name)

		return
	}

	// protocol: https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ipheader := ipHeader{
		version:        packet[0] >> 4,      // 先頭の byte の 4 bit
		headerLen:      packet[0] << 5 >> 5, // 先頭の byte の後半 4 bit
		tos:            packet[1],           // 1 byte
		totalLen:       byteToUint16(packet[2:4]),
		identify:       byteToUint16(packet[4:6]),
		fragOffset:     byteToUint16(packet[6:8]),
		ttl:            packet[8],
		protocol:       packet[9], // TCP: 6, ICMP: 1
		headerChecksum: byteToUint16(packet[10:12]),
		srcAddr:        byteToUint32(packet[12:16]),
		destAddr:       byteToUint32(packet[16:20]),
	}

	fmt.Printf("ipInput Received IP in %s, packet type %d from %s to %s\n", inputdev.name, ipheader.protocol,
		printIPAddr(ipheader.srcAddr), printIPAddr(ipheader.destAddr))

	// IPバージョンが4でなければドロップ
	// Todo: IPv6の実装
	if ipheader.version != 4 {
		if ipheader.version == 6 {
			fmt.Println("packet is IPv6")
		} else {
			fmt.Println("Incorrect IP version")
		}
		return
	}

	// IPヘッダオプションがついていたらドロップ = ヘッダ長が20byte以上だったら
	if 20 < (ipheader.headerLen * 4) {
		fmt.Println("IP header option is not supported")
		return
	}

	// 宛先アドレスがブロードキャストアドレス or 受信した NIC インターフェイスのIPアドレスの場合
	if (ipheader.destAddr == IP_ADDRESS_LIMITED_BROADCAST) ||
		(inputdev.ipdev.address == ipheader.destAddr) {
		// 自分宛の通信として処理。
		ipInputToOurs(inputdev, &ipheader, packet[20:])

		return
	}

	// 宛先 IP アドレスをルータが持ってるか調べる。
	for _, dev := range netDeviceList {
		// 宛先IPアドレスがルータの持っているIPアドレス or ディレクティッド・ブロードキャストアドレスの時の処理
		if dev.ipdev.address == ipheader.destAddr || dev.ipdev.broadcast == ipheader.destAddr {
			// 自分宛の通信として処理
			ipInputToOurs(inputdev, &ipheader, packet[20:])

			return
		}
	}

	// TTLが1以下ならドロップ
	if ipheader.ttl <= 1 {
		// Todo: send_icmp_time_exceeded関数を作成
		return
	}

	// TTLを1へらす
	ipheader.ttl -= 1

	// IPヘッダチェックサムの再計算
	ipheader.headerChecksum = 0
	ipheader.headerChecksum = byteToUint16(calcChecksum(ipheader.ToPacket(true)))

	// my_buf構造にコピー
	// forwardPacket := ipheader.ToPacket(true)
	// NATの内側から外側への通信
	// if inputdev.ipdev.natdev != (natDevice{}) {
	// 	forwardPacket = append(forwardPacket, natPacket...)
	// } else {
	// 	forwardPacket = append(forwardPacket, packet[20:]...)
	// }

	// if route.iptype == connected { // 直接接続ネットワークの経路なら
	// 	// hostに直接送信
	// 	ipPacketOutputToHost(route.netdev, ipheader.destAddr, forwardPacket)
	// } else { // 直接接続ネットワークの経路ではなかったら
	// 	fmt.Printf("next hop is %s\n", printIPAddr(route.nexthop))
	// 	fmt.Printf("forward packet is %x : %x\n", forwardPacket[0:20], natPacket)
	// 	ipPacketOutputToNetxhop(route.nexthop, forwardPacket)
	// }
}

/*
自分宛のIPパケットの処理
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/ip.cpp#L26
*/
func ipInputToOurs(inputdev *netDevice, ipheader *ipHeader, packet []byte) {
	// 上位プロトコルの処理に移行する。
	switch ipheader.protocol {
	case IP_PROTOCOL_NUM_ICMP:
		fmt.Println("ICMP received!")
		icmpInput(inputdev, ipheader.srcAddr, ipheader.destAddr, packet)
	case IP_PROTOCOL_NUM_UDP:
		fmt.Printf("udp received : %x\n", packet)

		return
	case IP_PROTOCOL_NUM_TCP:
		return
	default:
		fmt.Printf("Unhandled ip protocol number : %d\n", ipheader.protocol)

		return
	}
}

type icmpHeader struct {
	icmpType uint8
	icmpCode uint8
	checksum uint16
}

type icmpEcho struct {
	identify  uint16
	sequence  uint16
	timestamp []uint8
	data      []uint8
}

type icmpDestinationUnreachable struct {
	unused uint32
	data   []uint8
}

type icmpTimeExceeded struct {
	unused uint32
	data   []uint8
}

type icmpMessage struct {
	icmpHeader                 icmpHeader
	icmpEcho                   icmpEcho
	icmpDestinationUnreachable icmpDestinationUnreachable
	icmpTimeExceeded           icmpTimeExceeded
}

const (
	ICMP_TYPE_ECHO_REPLY              uint8 = 0
	ICMP_TYPE_DESTINATION_UNREACHABLE uint8 = 3
	ICMP_TYPE_ECHO_REQUEST            uint8 = 8
	ICMP_TYPE_TIME_EXCEEDED           uint8 = 11
)

func (icmpmsg icmpMessage) ReplyPacket() (icmpPacket []byte) {
	var b bytes.Buffer
	// ICMPヘッダ
	b.Write([]byte{ICMP_TYPE_ECHO_REPLY})
	b.Write([]byte{0x00})       // icmp code
	b.Write([]byte{0x00, 0x00}) // checksum
	// ICMPエコーメッセージ
	b.Write(uint16ToByte(icmpmsg.icmpEcho.identify))
	b.Write(uint16ToByte(icmpmsg.icmpEcho.sequence))
	b.Write(icmpmsg.icmpEcho.timestamp)
	b.Write(icmpmsg.icmpEcho.data)

	icmpPacket = b.Bytes()
	checksum := calcChecksum(icmpPacket)
	// 計算したチェックサムをセット
	icmpPacket[2] = checksum[0]
	icmpPacket[3] = checksum[1]

	fmt.Printf("Send ICMP Packet is %x\n", icmpPacket)

	return icmpPacket
}

func icmpInput(inputdev *netDevice, sourceAddr, destAddr uint32, icmpPacket []byte) {
	// ICMP メッセージ長より短い場合。
	if len(icmpPacket) < 4 {
		fmt.Println("Received ICMP Packet is too short")
	}

	// ICMPのパケットとして解釈する
	icmpmsg := icmpMessage{
		icmpHeader: icmpHeader{
			icmpType: icmpPacket[0],
			icmpCode: icmpPacket[1],
			checksum: byteToUint16(icmpPacket[2:4]),
		},
		icmpEcho: icmpEcho{
			identify:  byteToUint16(icmpPacket[4:6]),
			sequence:  byteToUint16(icmpPacket[6:8]),
			timestamp: icmpPacket[8:16],
			data:      icmpPacket[16:],
		},
	}

	switch icmpmsg.icmpHeader.icmpType {
	case ICMP_TYPE_ECHO_REPLY:
		fmt.Println("ICMP ECHO REPLY is received")
	case ICMP_TYPE_ECHO_REQUEST:
		fmt.Println("ICMP ECHO REQUEST is received, Create Reply Packet")
		ipPacketEncapsulateOutput(inputdev, sourceAddr, destAddr, icmpmsg.ReplyPacket(), IP_PROTOCOL_NUM_ICMP)
	}
}

/*
IPパケットにカプセル化して送信する。
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/ip.cpp#L102
*/
func ipPacketEncapsulateOutput(inputdev *netDevice, destAddr, srcAddr uint32, payload []byte, protocolType uint8) {
	// IPヘッダで必要なIPパケットの全長を算出する。
	// IPヘッダの20byte + パケットの長さ。
	totalLength := 20 + len(payload)

	ipheader := ipHeader{
		version:        4, // IPv4
		headerLen:      20 / 4,
		tos:            0,
		totalLen:       uint16(totalLength),
		identify:       0xf80c,
		fragOffset:     2 << 13,
		ttl:            0x40,
		protocol:       protocolType,
		headerChecksum: 0, // checksum計算する前は0をセット
		srcAddr:        srcAddr,
		destAddr:       destAddr,
	}

	var ipPacket []byte
	ipPacket = append(ipPacket, ipheader.ToPacket(true)...)
	// payload を、IP ヘッダの後ろに追加
	ipPacket = append(ipPacket, payload...)

	// ルートテーブルを検索して送信先IPのMACアドレスがなければ、
	// ARPリクエストを生成して送信して結果を受信してから、ethernetからパケットを送る
	destMacAddr, _ := searchArpTableEntry(destAddr)
	if destMacAddr != [6]uint8{0, 0, 0, 0, 0, 0} {
		// ルートテーブルに送信するIPアドレスのMACアドレスがあれば送信
		ethernetOutput(inputdev, destMacAddr, ipPacket, ETHER_TYPE_IP)
	} else {
		// ARPリクエストを出す
		sendArpRequest(inputdev, destAddr)
	}
}

/*
ARPテーブルの検索
*/
func searchArpTableEntry(ipaddr uint32) ([6]uint8, *netDevice) {
	if len(ArpTableEntryList) != 0 {
		for _, arpTable := range ArpTableEntryList {
			if arpTable.ipAddr == ipaddr {
				return arpTable.macAddr, arpTable.netdev
			}
		}
	}
	return [6]uint8{}, nil
}

/*
ARPリクエストの送信
https://github.com/kametan0730/interface_2022_11/blob/master/chapter2/arp.cpp#L111
*/
func sendArpRequest(netdev *netDevice, targetip uint32) {
	fmt.Printf("Sending arp request via %s for %x\n", netdev.name, targetip)
	// APRリクエストのパケットを作成
	arpPacket := arpIPToEthernet{
		hardwareType:        ARP_HTYPE_ETHERNET,
		protocolType:        ETHER_TYPE_IP,
		hardwareLen:         ETHERNET_ADDRES_LEN,
		protocolLen:         IP_ADDRESS_LEN,
		opcode:              ARP_OPERATION_CODE_REQUEST,
		senderHardwareAddr:  netdev.macaddr,
		senderIPAddr:        netdev.ipdev.address,
		targetHardwareAddrr: ETHERNET_ADDRESS_BROADCAST,
		targetIPAddr:        targetip,
	}.ToPacket()
	// ethernetでカプセル化して, FF-FF-FF-FF-FF-FF を送信する。
	ethernetOutput(netdev, ETHERNET_ADDRESS_BROADCAST, arpPacket, ETHER_TYPE_ARP)
}

func getIPdevice(addrs []net.Addr) (ipdev ipDevice) {
	for _, addr := range addrs {
		// ipv6ではなくipv4アドレスをリターン
		ipaddrstr := addr.String()
		if !strings.Contains(ipaddrstr, ":") && strings.Contains(ipaddrstr, ".") {
			ip, ipnet, _ := net.ParseCIDR(ipaddrstr)
			ipdev.address = byteToUint32(ip.To4())
			ipdev.netmask = byteToUint32(ipnet.Mask)
			// ブロードキャストアドレスの計算はIPアドレスとサブネットマスクのbit反転の2進数「OR（論理和）」演算
			ipdev.broadcast = ipdev.address | (^ipdev.netmask)
		}
	}
	return ipdev
}
