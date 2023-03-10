package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	computer = strings.NewReader("COMPUTER")
	system   = strings.NewReader("SYSTEM")
	karin    = strings.NewReader("KARIN")
)

// Practice about Concurrent & Parallel
//
// Especially, using these functions
//
//	goroutines
func main() {
	var stream io.Reader
	stream = io.MultiReader(computer, system, karin)
	io.Copy(os.Stdout, stream)

	// calcLoan()
	var builder strings.Builder
	builder.Write([]byte("strings.Builder write\n"))
	fmt.Println(builder.String())

	f, _ := os.Create("test")
	// 1024 バイト（= 1 KB）。
	buffer := make([]byte, 1024)
	// io.Copy では 32 KB のバッファが内部で確保されている。
	io.CopyBuffer(f, f, buffer)

	var reader io.Reader = strings.NewReader("test data")
	// ダミーの Close() メソッドを持って io.ReadCloser のふりをする！
	var readCloser io.ReadCloser = ioutil.NopCloser(reader)
	defer readCloser.Close()

	file, _ := os.Open("main.go")
	defer file.Close()
	io.Copy(io.MultiWriter(f, os.Stdout), file)

	strings.NewReader("Reader へ文字列でわたせる")

	scanner := bufio.NewScanner(strings.NewReader("hoge\nfuga\npien"))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

	randFile()

	// ================= png =================
	createPNGWithText()
	file, err := os.Open("test-new.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func randFile() {
	f, err := os.Create("random")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 1024)
	rand.Reader.Read(buf)

	f.Write(buf)
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	// tEXt チャンクのみ中身を表示する。
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

// PNG はチャンクから構成され、チャンクは
// 長さ(4 bytes) + 種類(4 bytes) + データ(長さ) + CRC(4 bytes) からなる
func createPNGWithText() {
	file, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("test-new.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("KARIN-SAMA-kawaii"))
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	// 最初の 8 バイトは飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		offset, _ = file.Seek(int64(length+8), 1)
	}

	return chunks
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)
	// CRC を追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func calc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	fmt.Printf("year=%d total=%d interest=%d id=%d\n", year, price+interest, interest, id)
}
func worker(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
	// タスクがなくなってタスクのチャネルが close されるまで無限ループ
	for yaer := range years {
		calc(id, price, interestRate, yaer)
		wg.Done()
	}
}

func calcLoan() {
	// 借入額
	price := 4000_0000
	// 利子1.1%
	interestRate := 0.011
	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		years <- i
	}
	var wg sync.WaitGroup
	wg.Add(35)
	// CPU コア数分の goroutine 起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}
	close(years)
	wg.Wait()
}

func goroutineCost() {
	tasks := []string{
		"go build -o main main.go",
		"mv main share",
		"./publish",
	}
	for _, task := range tasks {
		go func() {
			// CAUTION!
			// goroutine is very fast, but not cost zero
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
