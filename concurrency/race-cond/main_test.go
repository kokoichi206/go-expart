package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "34320.00") {
		t.Error("error")
	}
}

// func Test_updateMessage(t *testing.T) {
// 	msg = "hi"

// 	wg.Add(1)
// 	go updateMessage("byebye")
// 	wg.Wait()

// 	if msg != "byebye" {
// 		t.Error("inccorect")
// 	}
// }
