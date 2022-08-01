package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Laps"] = Vertex{40.23454, -69.23454}
	fmt.Println(m["Bell Laps"])
	fmt.Println(m)
}
