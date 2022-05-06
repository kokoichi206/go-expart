package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "http://localhost:18888"

	var wg sync.WaitGroup
	wg.Add(35)

	N := 100
	for i := 0; i < N; i++ {
		fmt.Println(i)
		go dos(url, &wg)
	}

	wg.Wait()
}

func dos(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	wg.Done()
	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray))
}
