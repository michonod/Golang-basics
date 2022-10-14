package main

import "fmt"

func main() {
	//This is how you can define map
	emails := make(map[string]string)

	emails["Mihail"] = "mihail@hotmail.com"
	emails["Tara"] = "tara@hotmail.com"

	fmt.Println(emails)
	//delete from map
	delete(emails, "Tara")
	fmt.Println(emails)

	//declare map at first

	age := map[string]int{"Mihail": 30, "Tara": 29}

	fmt.Println(age)

}
