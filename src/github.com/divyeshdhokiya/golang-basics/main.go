package main

import "fmt"

const b int16 = 90

func main() {
	/* Const */
	const a int = 10
	const aa = 11                // implicit conversion
	fmt.Printf("%v, %T", a, a)   // 10, int
	fmt.Printf("%v, %T", aa, aa) // 11, int

	// shadow
	const b int = 89
	fmt.Printf("%v, %T", b, b) // 89, int
}
