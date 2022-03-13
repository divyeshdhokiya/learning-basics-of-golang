package main

import (
	"fmt"
)

func main() {
	/* Boolean */
	var n bool = false
	fmt.Printf("%v, %T", n, n)

	a := 1 == 1
	fmt.Printf("%v,%T", a, a)
	b := 12 == 13
	fmt.Printf("%v,%T", b, b)

}
