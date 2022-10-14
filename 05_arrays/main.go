package main

import "fmt"

func main() {

	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	appleColor := [2]string{"Red", "Blue"}
	var bananaColor = [3]string{"Red", "Green", "Yellow"}
	fruitSlice := []string{"Red", "Blue", "Red", "Blue", "Red", "Blue", "Red", "Blue"}

	fmt.Println(fruitArr, appleColor, bananaColor, len(fruitSlice))
	fmt.Println(bananaColor[1:3])
}
