package main

import (
	"fmt"
	"monkey-language/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %v\n", user)

	repl.Start(os.Stdin, os.Stdout)
}
