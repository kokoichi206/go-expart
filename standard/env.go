package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func envTest() {
	os.Setenv("STANDARD_GO", "standard")
	res, err := exec.Command("echo", "$PATH").CombinedOutput()
	if err != nil {
		panic(err)
	}
	os.Environ()

	fmt.Printf("string(res): %v\n", string(res))

	// Goのos/execパッケージのCommand関数は、シェルを経由せずに直接プログラムを実行します。
	// そのため、シェル特有の機能（例えば環境変数の展開）を使用することはできません。
	cmd := exec.Command("bash", "-c", "echo $PATH")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", output)

}
