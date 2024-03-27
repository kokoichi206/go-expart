package hat

import (
	"io"
	"log"
	"net"
	"os"
)

func handle(src net.Conn) {
	// HOST_PORT = "localhost:9000"
	host := os.Getenv("HOST_PORT")
	if host == "" {
		log.Fatal("HOST_PORT environment variable is not set.")
	}

	dst, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatalf("net.Dial: %v", err)
	}
	defer dst.Close()

	// Run in goroutine to prevent blocking
	// Copy is a blocking operation
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalf("io.Copy(dst, src): %v", err)
		}
	}()

	// Copy from destination to source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalf("io.Copy(src, dst): %v", err)
	}
}

func Proxy() {
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}
	defer l.Close()

	for {
		// Accept incoming connections
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("l.Accept: %v", err)
		}

		// Handle connections in a new goroutine
		go handle(conn)
	}
}
