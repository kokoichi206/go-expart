package main

import (
	"fmt"
	"slices"
)

type user struct {
	name string
	age  int
}

func main() {
	u := &user{
		name: "John",
		age:  30,
	}
	u2 := &user{
		name: "Jane",
		age:  25,
	}
	u3 := &user{
		name: "Doe",
		age:  40,
	}
	users := []*user{u, u2, u3}

	// slices.Clone
	p := slices.Clone(users)
	fmt.Printf("p: %v\n", p)
	p[0].name = "Alice"
	fmt.Printf("p: %v\n", p)
	fmt.Printf("users: %v\n", users)
	fmt.Printf("p[0]: %v\n", p[0])
	fmt.Printf("users[0]: %v\n", users[0])

	// 以前のやり方、これも結局 shallow copy
	q := append([]*user{}, users...)
	fmt.Printf("q: %v\n", q)
	q[0].name = "Bob"
	fmt.Printf("q: %v\n", q)
	fmt.Printf("users: %v\n", users)
	fmt.Printf("q[0]: %v\n", q[0])
	fmt.Printf("users[0]: %v\n", users[0])

	// build-in copy
	r := make([]*user, len(users))
	copy(r, users)
	r[0].name = "Charlie"
	fmt.Printf("r: %v\n", r)
	fmt.Printf("users: %v\n", users)
	fmt.Printf("r[0]: %v\n", r[0])
	fmt.Printf("users[0]: %v\n", users[0])

	s := make([]*user, len(users))
	for i, v := range users {
		// sync.Mutex などを使っている場合は、この方法は使えない。
		vv := *v
		s[i] = &vv
	}
	s[0].name = "David"
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s[0]: %v\n", s[0])
	fmt.Printf("users[0]: %v\n", users[0])
}
