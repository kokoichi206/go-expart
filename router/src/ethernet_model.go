package main

import "syscall"

type ethernetHeader struct {
	destAddr  [6]uint8
	srcAddr   [6]uint8
	etherType uint16
}

type netDevice struct {
	name       string
	macaddr    [6]uint8
	socket     int
	socketaddr syscall.SockaddrLinklayer
	etheHeader ethernetHeader
}
