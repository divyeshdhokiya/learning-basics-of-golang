package main

import (
	"fmt"
)

func main() {
	/* Looping */

	// Simple Loops
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
		/*
			OP:
				0 0
				1 2
				2 4
				3 6
				4 8
		*/
	}

	k := 0 // scope to main function
	for k < 5 {
		fmt.Println(k)
		k++
	}
	/*
		OP:
		0
		1
		2
		3
		4
	*/
	fmt.Println("--")

	// Exiting Early
	l := 0
	for {
		if l > 5 {
			break
		}
		fmt.Println(l)
		l++
	}
	/*
		OP:
		0
		1
		2
		3
		4
		5
	*/
	println("--")

	// break out from parent look
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	/*
		OP:
		1
		2
		3
	*/
	println("--")

	// Looping through collections - for range loop
	s := []int{1, 2, 3}
	for m, n := range s {
		fmt.Println(m, n)
	}
	/*
		OP:
		0 1
		1 2
		2 3
	*/
	// Ignore index
	for _, n := range s {
		fmt.Println(n)
	}
	/*
		OP:
		1
		2
		3
	*/
}
