package main

import "fmt"

func main() {
	ids := []int{33, 44, 55, 66, 77}
	fmt.Println("HI")

	//Loop through ids

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	emails := map[string]string{"Mihail": "mihail_davidevski@outlook.com", "Tara": "Tara@outlook.com"}

	//Range with map

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("%s\n", k)
	}

}
