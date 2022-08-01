package main

import "fmt"

func main() {
	total := 1
	for i := 1; i < 50; i++ {
		total *= i
		fmt.Println(total)
	}

	fmt.Println(total)
}
