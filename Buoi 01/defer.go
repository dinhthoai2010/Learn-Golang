package main

import "fmt"

func main() {
	// Trì hoãn chạy sau
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Tuy phong")

}
