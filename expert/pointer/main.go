package main

import "fmt"

func main() {
	// for i := 0; i < 1000; i++ {
	// 	NewHeroPtr()
	// }

	// v := pien{
	// 	paon:   "hogehogehogehoge",
	// 	hoge:   "fugafugafugafuga",
	// 	taihen: "taihen taihenda",
	// }

	// for i := 0; i < 1000; i++ {
	// 	_ = sg(&v)
	// }

	hi := []*pien{
		{
			paon:   "hogehogehogehoge",
			hoge:   "fugafugafugafuga",
			taihen: "taihen taihenda",
		},
	}

	for i := 0; i < 1000; i++ {
		_ = g(hi)
	}
}

func defaultCapSize() {
	// 	before: 0
	// before: 1
	// before: 2
	// before: 4
	// before: 8
	// before: 16
	// before: 32
	// before: 64
	// before: 128
	// before: 256
	// before: 512
	// before: 848
	// before: 1280
	testSlice := []string{}
	before := cap(testSlice)
	fmt.Printf("before: %v\n", before)
	for i := 0; i < 1000; i++ {
		if before != cap(testSlice) {
			before = cap(testSlice)
			fmt.Printf("before: %v\n", before)
		}
		testSlice = append(testSlice, "hoge")
	}
}
