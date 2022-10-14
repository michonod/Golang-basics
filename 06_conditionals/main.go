package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	color := "red"

	if color == "red" {
		fmt.Println("Color is simply RED!")
	} else if color == "green" {
		fmt.Println("Color is green!")
	} else {
		fmt.Println("Don't know this color...")
	}

	switch color {
	case "red":
		fmt.Println("This color is red!")
	case "blue":
		fmt.Println("This color is blue!")
	case "green":
		fmt.Println("This color is green!")
	default:
		fmt.Println("I don't know this color! ")
	}
}
