package main

import (
	"fmt"
	"strconv"
)

//Define vehicle struct

type Vehicle struct {
	brand string
	motor string
	color string
	kw    int
}

// Brooming
func (v Vehicle) broom(sound string) string {
	return sound + v.brand + strconv.Itoa(v.kw)
}

func main() {
	// Init vehicle

	hondaCivic := Vehicle{brand: "Honda", motor: "2.2 CDTI", color: "grey"}

	fmt.Println(hondaCivic)
	fmt.Println(hondaCivic.broom("brrrrrrr brrrr "))
}
