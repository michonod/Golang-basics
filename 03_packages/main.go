package main

import (
	"fmt"
	"math"

	"github.com/michonod/starting_with_go/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor((4.7)))
	fmt.Println(math.Ceil((7.4)))
	fmt.Println(math.Sqrt((7)))
	fmt.Println(strutil.Reverse("This should be reversed!"))
}
