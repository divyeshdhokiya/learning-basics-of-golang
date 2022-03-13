package main

import (
	"fmt"
)

func main() {
	/* "if" statement */

	if true {
		fmt.Println("inside true block")
	}

	colors := map[string]string{
		"Red":   "#123456",
		"White": "#fff",
		"Black": "#000",
	}
	fmt.Println(colors["Red"])

	if color, ok := colors["White"]; ok {
		fmt.Println(color, ok) // OP: #fff true
	}

	if color, ok := colors["Blue"]; ok {
		fmt.Println(color, ok) // Not executed
	}

	a := 20
	b := 40
	if a < b {
		fmt.Println("Low")
	}
	if a > b {
		fmt.Println("High")
	}

}
