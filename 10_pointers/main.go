package main

import "fmt"

func main() {
	//memory address assign
	a := 5
	b := &a

	fmt.Println(a, b)

	//value of b
	fmt.Println(*b)

	//change val with pointer

	*b = 10
	fmt.Println(a, b)
}
