package main

import "fmt"

func main() {
	// Main types in GO
	// ---> string
	// ---> bool
	// ---> int
	// ---> int int8 int16 int32 int64
	// ---> uint uint8 uint16 uint32 uint64 uintptr
	// ---> byte - alias for uint8
	// ---> rune - alias for int32
	// ---> float32 float64
	// ---> complex64 complex128

	fmt.Println("Hello from GO!")

	// VAR
	// var name = "Mihail"
	//Shorthand for variable
	// name := "Mihail"
	// email := "mihail_davidevski@outlook.com"
	var age = 27
	var isCool = true
	var size float32 = 3.22

	name, email := "Mihail", "mihail_davidevski@outlook.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", age)

}
