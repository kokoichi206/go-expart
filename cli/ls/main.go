package main

import (
	"fmt"
	"io/ioutil"
	"log"

	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var target string
	if len(args) == 0 {
		target = "."
	} else if len(args) == 1 {
		target = args[0]
	} else {
		// TODO: Multiple arguments output
		return
	}
	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%s\t", file.Name())
	}
	fmt.Println()
}
