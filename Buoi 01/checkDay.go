package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When `s Saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Ngay mai")
	case today + 2:
		fmt.Println("Ngay kia")
	default:
		fmt.Println("Too far day")
	}
}
