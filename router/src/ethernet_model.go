package main

import (
	"bytes"
	"syscall"
)

type ethernetHeader struct {
	destAddr  [6]uint8
	srcAddr   [6]uint8
	etherType uint16
}

func (ethHeader ethernetHeader) ToPacket() []byte {
	var b bytes.Buffer
	b.Write(macToByte(ethHeader.destAddr))
	b.Write(macToByte(ethHeader.srcAddr))
	b.Write(uint16ToByte(ethHeader.etherType))
	return b.Bytes()
}

type netDevice struct {
	name       string
	macaddr    [6]uint8
	socket     int
	socketaddr syscall.SockaddrLinklayer
	etheHeader ethernetHeader
	ipdev      ipDevice
}

type ipDevice struct {
	address   uint32 // デバイスのIPアドレス
	netmask   uint32 // サブネットマスク
	broadcast uint32 // ブロードキャストアドレス
}
