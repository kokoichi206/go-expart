package main

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	// writerExample()
	// fileWrite()
	zipWrite()
}

func endian() {
	// 32 ビットのビックエンディアンのデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
}

func fileWrite() {
	file, _ := os.Create("test.txt")
	defer file.Close()
	file.Write([]byte("new line"))
	file.Write([]byte("new line"))
	file.Write([]byte("new line"))
	io.Copy(os.Stdout, file)
	// io.Copy(file, file)
}

func stdin() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}

func readExample() {
	// ラップする
	var reader io.Reader = strings.NewReader("test data")
	var readCloser io.ReadCloser = ioutil.NopCloser(reader)
	readCloser.Close()
}

func csvWrite() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, err := os.Create("test.tsv")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(file)
	w.Comma = '\t'

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func zipWrite() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Flush()
}

func internetAccess() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// net.Conn は、io.Reader のインタフェースでもあることを利用
	io.Copy(os.Stdout, conn)

	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	req.Write(conn)
	// http.ResponseWriter というものも高レイヤーにある
}

func writerExample() {
	fmt.Println("Hello World!")

	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))

	os.Stdout.Write([]byte("os.Stdout exapmle\n"))

	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
