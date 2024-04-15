package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func parse() {
	exp, err := parser.ParseExpr("v + 1")
	if err != nil {
		log.Fatal(err)
	}
	ast.Print(nil, exp)
}

func fileParse() {
	fmt.Println("----- fileParse -----")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "_yoshi.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// 抽象構文木の深さ優先探索！
	ast.Inspect(f, func(n ast.Node) bool {
		// 戻り値が false の場合 ast.Inspect は再帰を行わない。
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		// package ast // import "go/ast"

		// type Ident struct {
		//         NamePos token.Pos // identifier position
		//         Name    string    // identifier name
		//         Obj     *Object   // denoted object, or nil. Deprecated: see Object.
		// }
		//     An Ident node represents an identifier.

		// func NewIdent(name string) *Ident
		// func (x *Ident) End() token.Pos
		// func (id *Ident) IsExported() bool
		// func (x *Ident) Pos() token.Pos
		// func (id *Ident) String() string

		// fmt.Printf("ident.Name: %v\n", ident.Name)
		if ident.Name != "Position" {
			return true
		}

		// fmt.Printf("ident.Pos(): %v\n", ident.Pos())
		fmt.Println(fset.Position(ident.Pos()))

		return true
	})
}

// fileParse からの発展。
func typeCheck() {
	fmt.Println("----- typeCheck -----")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "_yoshi.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// 識別子が定義 or 利用されてる部分を記録するため。
	defsOrUses := map[*ast.Ident]types.Object{}
	// https://pkg.go.dev/go/types#Info
	info := &types.Info{
		Defs: defsOrUses,
		Uses: defsOrUses,
	}

	// 型チェックのための設定！
	config := &types.Config{
		// ここだけは must で必要。
		Importer: importer.Default(),
	}

	// 1: パッケージ名
	// 2: 字句解析時に記録したトークンの位置を保持した *token.FileSet
	// 3: ファイル単位の抽象構文木のスライス
	// 4: 型情報を格納するための *types.Info
	if _, err := config.Check("main", fset, []*ast.File{f}, info); err != nil {
		log.Fatal(err)
	}

	// 抽象構文木の深さ優先探索！
	ast.Inspect(f, func(n ast.Node) bool {
		// 戻り値が false の場合 ast.Inspect は再帰を行わない。
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		if ident.Name != "Position" {
			obj := defsOrUses[ident]
			if obj == nil {
				return true
			}
			typ := obj.Type()
			if _, ok := typ.(*types.Named); !ok {
				return true
			}
			if obj.Name() == "Age" {
				fmt.Printf("obj.Name(): %v\n", obj.Name())
				// 基底型は int であることがわかる。
				fmt.Printf("typ.Underlying(): %v\n", typ.Underlying())
			}
			return true
		}

		obj := defsOrUses[ident]
		if obj == nil {
			return true
		}

		// 型情報の取得。
		typ := obj.Type()
		if _, ok := typ.(*types.Named); !ok {
			fmt.Printf("typ: %v\n", typ)
			return true
		}

		fmt.Println(fset.Position(ident.Pos()))

		return true
	})
}

func main() {
	fileParse()

	typeCheck()
}
