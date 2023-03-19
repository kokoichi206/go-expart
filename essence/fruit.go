package main

type Fruit int

// このコマンドを書いておくと、go generate で生成できる！
//go:generate stringer -type Fruit fruit.go
const (
	Apple Fruit = iota
	Orange
	Banana
)
