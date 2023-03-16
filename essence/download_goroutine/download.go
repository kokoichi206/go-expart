package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV: ", err)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// 良くわかってない、普段 err の時 close 入ってるっけ？
			// （エラーハンドリングの後に defer で書いてしまってるイメージがあった）
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b
	}
}

func main() {
	urls := []string {
		"https://testhogehoge.piyon/xyz.csv",
		"https://testhogehoge.piyon/xyz2.csv",
		"https://testhogehoge.piyon/xyz3.csv",
	}

	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	// goroutine からコンテンツを受け取る！
	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			// レコードの db 登録
			fmt.Printf("records: %v\n", records)
		}
	}
	wg.Wait()
}
