package main

import "fmt"

func main() {
	fmt.Println("Hi")

	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 2; i <= 20; i++ {
		fmt.Println(i)
	}

	for f := 0; f <= 100; f++ {
		if f%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if f%3 == 0 {
			fmt.Println("Fizz")
		} else if f%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(f)
		}
	}
}
