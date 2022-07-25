package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

type Params struct {
	IsHelp     bool
	IsColor    bool
	ShowHidden bool
	Args       []string
	Name       string
}

// Print usage
var Usage = func() {
	fmt.Println("NAME")
	fmt.Printf("\tls – list directory contents\n\n")
	fmt.Printf("The following options are available:\n")
	flag.PrintDefaults()
}

// Is this global var ok?
var params Params
var osExit = os.Exit

func init() {
	flag.BoolVarP(&params.IsHelp, "help", "h", false, "Print help message")
	flag.BoolVarP(&params.IsColor, "G", "G", false, "Print with bold cyan")
	flag.BoolVarP(&params.ShowHidden, "all", "a", false, "Print hidden files")

	flag.Parse()

	params.Args = flag.Args()
}

func main() {

	output(params)
}

func output(params Params) {
	if params.IsHelp {
		Usage()

		// os.Exit and return are redundant... but for testing
		osExit(0)
		return
	}

	if len(params.Args) == 0 {
		params.Name = "."
	} else if len(params.Args) == 1 {
		params.Name = params.Args[0]
	} else {
		// TODO: Multiple arguments output
		osExit(0)
		return
	}

	files, err := ioutil.ReadDir(params.Name)
	if err != nil {
		fmt.Printf("ls: %s: No such file or directory\n", params.Name)
		osExit(1)
		return
	}

	if params.IsColor {
		printFilesWithColor(files, params.ShowHidden)
	} else {
		printFiles(files, params.ShowHidden)
	}

	fmt.Println()
}

func printFiles(files []fs.FileInfo, showHidden bool) {
	for _, file := range files {
		if !isHiddenFile(file) || showHidden {
			fmt.Printf("%s\t", file.Name())
		}
	}
}

func printFilesWithColor(files []fs.FileInfo, showHidden bool) {
	for _, file := range files {

		if !isHiddenFile(file) || showHidden {
			if file.IsDir() {
				d := color.New(color.FgCyan, color.Bold)
				d.Printf("%s\t", file.Name())
			} else {
				fmt.Printf("%s\t", file.Name())
			}
		}
	}
}

// Check whether the file is hidden file or not.
func isHiddenFile(file fs.FileInfo) bool {

	// TODO: Need more check especially for Windows or...
	return isStartDot(file)
}

func isStartDot(file fs.FileInfo) bool {

	const dotCharacter = 46

	return len(file.Name()) > 0 && file.Name()[0] == dotCharacter
}
