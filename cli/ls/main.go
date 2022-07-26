package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

type Params struct {
	IsHelp     bool
	IsColor    bool
	ShowHidden bool
	ShowList   bool
	Args       []string
	Name       string
}

const APP_NAME = "ls"

// Print usage
var Usage = func() {
	fmt.Println("NAME")
	fmt.Printf("\t%s – list directory contents\n\n", APP_NAME)
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
	flag.BoolVarP(&params.ShowList, "list", "l", false, "Print hidden files")

	flag.Parse()

	params.Args = flag.Args()
}

func main() {

	execute(params)
}

func execute(params Params) {
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

	printFiles(files, params.ShowHidden, params.ShowList, params.IsColor)

	fmt.Println()
}

func printFiles(files []fs.FileInfo, showHidden bool, showLong bool, showColor bool) {

	maxDigit := findMaxDigit(files)

	for _, file := range files {

		// 隠しファイルで、-a オプションがない場合は出力しない。
		if isHiddenFile(file) && !showHidden {
			continue
		}

		// -l オプションがあればそれ用のを出力する。
		if showLong {
			printLongInfo(file, maxDigit, showColor)
			continue
		}

		// 標準の出力。
		printBasic(file, showColor)
	}
}

// Check whether the file is hidden file or not.
func isHiddenFile(file fs.FileInfo) bool {

	// TODO: Need more check especially for Windows or...
	return isStartDot(file)
}

// ファイル名の開始が "." かどうか。
func isStartDot(file fs.FileInfo) bool {

	const dotCharacter = 46

	return len(file.Name()) > 0 && file.Name()[0] == dotCharacter
}

func printBasic(file fs.FileInfo, showColor bool) {
	if file.IsDir() && showColor {
		d := color.New(color.FgCyan, color.Bold)
		d.Printf("%s\t", file.Name())
	} else {
		fmt.Printf("%s\t", file.Name())
	}
}

// -l オプション指定時の出力。
func printLongInfo(file fs.FileInfo, maxDigit int, showColor bool) {
	p := file.Sys()
	var owner, group string
	nl := 0
	if statt, ok := p.(*syscall.Stat_t); ok {
		nl = int(statt.Nlink)
		uid := strconv.Itoa(int(statt.Uid))
		u, err := user.LookupId(uid)
		if err != nil {
			owner = uid
		} else {
			owner = u.Username
		}
		gid := strconv.Itoa(int(statt.Gid))
		g, err := user.LookupGroupId(gid)
		if err != nil {
			group = uid
		} else {
			group = g.Name
		}
	}
	const tf = "2006/01/02 15:04"
	ftime := fmt.Sprintf(file.ModTime().Format(tf))

	fSize := strconv.Itoa(int(file.Size()))

	var sb strings.Builder
	for i := 0; i < maxDigit-len(fSize); i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(fSize)

	var fname string
	if showColor && file.IsDir() {
		d := color.New(color.FgCyan, color.Bold)
		fname = d.Sprintf(file.Name())
	} else {
		fname = file.Name()
	}

	fmt.Printf("%s  %d %s %s %s %s %s\n", file.Mode(), nl, owner, group, sb.String(), ftime, fname)
}

func findMaxDigit(files []fs.FileInfo) int {
	max := 0
	for _, file := range files {

		digit := int(len(strconv.Itoa(int(file.Size()))))
		if digit > max {
			max = digit
		}
	}
	return max
}
