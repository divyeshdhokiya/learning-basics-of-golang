package main

import "fmt"

func main() {
	/* Defer: invoke a function but delay it's execution   */
	// c := 10
	// d := 20
	// defer fmt.Println(c, d)
	// e := c + d
	// fmt.Println(c, d, e)

	/*
		OP:
		10 20 30
		10 20
	*/

	// Order: Normal
	// fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	/*
		OP:
		1
		3
		2
	*/

	// Order: LIFO - Last in first out
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	/*
		OP:
		3
		2
		1
	*/

}
