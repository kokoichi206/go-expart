package main

import "strings"

type Point struct {
	X float64
	Y float64
}

type Node struct {
	// 左下と右上の2点。
	Min *Point
	Max *Point

	// 根からの深さ。
	Depth int32
}

func (n *Node) Children() []*Node {
	dx := (n.Max.X - n.Min.X) / 2.0
	dy := (n.Max.Y - n.Min.Y) / 2.0

	children := make([]*Node, 4)
	for idx := range children {
		ch := &Node{
			Min: &Point{
				X: n.Min.X,
				Y: n.Min.Y,
			},
			Max: &Point{
				X: n.Min.X + dx,
				Y: n.Min.Y + dy,
			},
			Depth: n.Depth + 1,
		}

		// 1 or 3
		if (idx & (1 << 0)) != 0 {
			ch.Min.X += dx
			ch.Max.X += dx
		}

		// 2 or 3
		if (idx & (1 << 1)) != 0 {
			ch.Min.Y += dy
			ch.Max.Y += dy
		}

		children[idx] = ch
	}

	return children
}
func (n *Node) IsInside(p *Point) bool {
	return n.Min.X <= p.X && p.X <= n.Max.X && n.Min.Y <= p.Y && p.Y <= n.Max.Y
}

type Tree struct {
	min *Point
	max *Point
}

func (t *Tree) Path(p *Point, depth int32) (*Node, string) {
	node := &Node{
		Min: t.min,
		Max: t.max,
	}

	builder := &strings.Builder{}
	label := "0123"

	for node.Depth < depth {
		for idx, ch := range node.Children() {
			if ch.IsInside(p) {
				node = ch
				builder.WriteByte(label[idx])

				break
			}
		}
	}

	return node, builder.String()
}

func main() {
}
