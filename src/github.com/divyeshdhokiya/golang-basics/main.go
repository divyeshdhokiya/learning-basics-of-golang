package main

import "fmt"

func main() {
	/* Complex */
	var n complex64 = 1 + 2i
	var n128 complex128 = 1 + 2i
	fmt.Printf("%v,%T\n", n, n)       // OP: (1+2i),complex64
	fmt.Printf("%v,%T\n", n128, n128) // OP: (1+2i),complex128

}
