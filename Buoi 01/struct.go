package main

import "fmt"

type Vecter struct {
	X int
	Y int
}

func Pointers() Vecter {
	p := Vecter{10, 20}
	q := &p
	q.X = 1e9

	return p
}

func main() {
	a := Vecter{1, 2}
	b := a
	b.X = 10
	fmt.Println(a.X, a.Y, b.X)

	fmt.Println(Pointers())
}
