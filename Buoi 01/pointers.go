package main

import "fmt"

func main() {
	// Khai báo con trỏ trong GO
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)

}
