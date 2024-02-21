package main

import (
	"fmt"
	"sync"
)

type store struct {
	data []string
}

func (s *store) add(wg *sync.WaitGroup, d string) {
	defer wg.Done()
	s.data = append(s.data, d)
}

func panipani() {
	panic("pien")
}

func testFunc() {
	panipani()
}

func main() {
	s := store{}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go s.add(wg, "hi")
	go s.add(wg, "hello")

	wg.Wait()

	fmt.Printf("s.data: %v\n", s.data)

	testFunc()
}
