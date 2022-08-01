package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go run on")
	switch os := runtime.GOOS; os {
	case "drawin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. \n ", os)
	}
}
