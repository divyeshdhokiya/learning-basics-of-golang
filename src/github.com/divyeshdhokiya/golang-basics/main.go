package main

import "fmt"

func main() {
	/* Numeric */
	// Signed Interger
	n := 32
	fmt.Printf("%v,%T", n, n)

	// Unsigned Integer
	var un uint16 = 42
	fmt.Printf("%v, %T", un, un)

	// bit operators
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010 -> 2
	fmt.Println(a | b)  // 1011 -> 11
	fmt.Println(a ^ b)  // 1001 -> 9
	fmt.Println(a &^ b) // 0100 ->  8

	// bit shifting
	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 -> 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 -> 1

	
}
