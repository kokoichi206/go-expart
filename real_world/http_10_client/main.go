package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
)

func main() {
	postMultipart()
}

func postMultipart() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "John Doe")
	// fileWriter, err := writer.CreateFormFile("parrot", "form_photo.png")

	readFile, err := os.Open("../imgs/parrot_sleep.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	// Content-Type 等を設定する
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/png")
	part.Set("Content-Disposition", `form-data; name="parrot" filename="form_photo.png"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888?%s", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}

func post() {
	// io.Reader を body に渡す
	// reader := strings.NewReader("my text")
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888?%s", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}

func get() {
	values := url.Values{
		"query": {"hello world"},
		"key":   {"must be in bracket"},
	}
	resp, err := http.Get(fmt.Sprintf("http://localhost:18888?%s", values.Encode()))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Printf("Headers: %v", resp.Header)
}
