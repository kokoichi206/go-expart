package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	All  bool `short:"a" long:"all" description:"Show all (including hidden items)"`
	List bool `short:"l" long:"list" description:"Show list information"`

	// Example of a callback, called each time the option is found.
	Call func(string) `short:"c" description:"Call phone number"`
}

func main() {
	opts.Call = func(num string) {
		cmd := exec.Command("open", "https://github.com/jessevdk/go-flags:"+num)
		cmd.Start()
		cmd.Process.Release()
	}

	args, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		panic(err)
	}

	fmt.Printf("All: %v\n", opts.All)
	fmt.Printf("List: %v\n", opts.List)
	fmt.Printf("Remaining args: %s\n", strings.Join(args, " "))

	commands := []string{}
	if opts.All {
		commands = append(commands, "-a")
	}
	if opts.List {
		commands = append(commands, "-l")
	}

	// exec で外部コマンドを使用する。
	ls, err := exec.Command("./ls", commands...).Output()
	if err != nil {
		fmt.Printf("./ls result:\n%s", ls)
		return
	}
	fmt.Println(ls)
}
