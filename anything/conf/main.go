//go:debug
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"slices"
)

type Validator interface{ Validate() error }

func JSONDecode(r io.Reader, x any) error {
	if err := json.NewDecoder(r).Decode(x); err != nil {
		return err
	}
	if v, ok := x.(Validator); ok {
		return v.Validate()
	}
	return nil
}

type U struct {
}

// 型制約。
func Print[T fmt.Stringer](s []T) {
	for _, v := range s {
		fmt.Printf("0x%s\n", v.String())
	}
}

func main() {
	// 構造体の中の順番を変え、よくアクセスされるものを1つ目に持ってきた。
	// メモリのアクセスが、番地 + offset でアクセスされる。
	// 関数展開すると 1 の時にのみ inline 展開されたり。
	// defer で increment するのは panic した時にも呼ばれるため。
	// o := sync.Once{}

	u := &U{}
	fmt.Printf("%[1]T %[1]v\n", u)

	sl := []string{"hoge", "pien"}
	// pa := []string{}
	pa := slices.Clone(sl)
	pa[0] = "fuga"
	fmt.Printf("sl: %v\n", sl)
	fmt.Printf("pa: %v\n", pa)

	m := map[string]string{"hoge": "pien"}
	cm := map[string]string{}
	// This is shallow copy
	// sm := maps.Clone(m)
	maps.Copy(cm, m)
	cm["hoge"] = "fuga"
	fmt.Printf("m: %v\n", m)
	fmt.Printf("cm: %v\n", cm)

	// clear は built-in になった！
	clear(cm)
	fmt.Printf("cm: %v\n", cm)
}
