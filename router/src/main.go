package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "ch1", "set run router mode")
	flag.Parse()

	if mode == "ch1" {
		runChap1()
	} else {
		fmt.Println("unexpected mode was passed.")
	}
}

func runChap1() {
	var netDeviceList []netDevice

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
		netDeviceList = append(netDeviceList, netDevice{
			name:       netif.Name,
			macaddr:    setMacAddr(netif.HardwareAddr),
			socket:     sock,
			socketaddr: addr,
		})
	}

	for {
		// パケットの受信を待つ！
		nfds, err := syscall.EpollWait(epfd, events, -1)
		if err != nil {
			log.Fatalf("failed to epoll wait: %s", err)
		}

		for i := 0; i < nfds; i++ {
			for _, netdev := range netDeviceList {
				// イベントがあったソケットとマッチしたら、
				// パケットを読み込む処理を実行！
				if events[i].Fd == int32(netdev.socket) {
					if err := netdev.netDevicePoll("ch1"); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
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
	}

	return nil
}
