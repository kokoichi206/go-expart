package main

import "fmt"

type Age int

type Me struct {
	// この辺は識別子としての Position とは違うので弾きたい。
	Position int
}

type Position struct {
	Y int
	X int
}

func NewPosition() Position {
	return Position{
		X: 1,
		Y: 2,
	}
}

func run() {
	const position = "1:2"
	var pos Position

	pos = NewPosition()

	fmt.Println("%v\n", pos)

	m := Me{
		Position: 1,
	}
	fmt.Println(m.Position)
}
