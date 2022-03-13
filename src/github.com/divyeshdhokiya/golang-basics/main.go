package main

import (
	"fmt"
)

func main() {
	/* "switch" statement */

	switch 2 {
	case 1, 2:
		fmt.Println("First case")
	case 3:
		fmt.Println("Second case")
	default:
		fmt.Println("Second case")

	}

	/* Tagless syntax */
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than")
	case i <= 20:
		fmt.Println("Less than")

	}

}
