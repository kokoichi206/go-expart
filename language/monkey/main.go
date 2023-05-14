package main

import (
	"fmt"
	"monkey-language/repl"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) > 1 {
		// input file.
		repl.StartFile(os.Args[1])
	} else {
		// interactive mode (REPL).
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)

		repl.Start(os.Stdin, os.Stdout)
	}
}
