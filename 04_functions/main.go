package main

import "fmt"

func greeting(name string) string {
	return "Hello my name is " + name
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Mihail"))
	fmt.Println(sum(2, 5))
}
