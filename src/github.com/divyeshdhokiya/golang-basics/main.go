package main

import (
	"fmt"
)

var I int = 10 // First letter Uppercase: Export variable
var i int = 20 // First letter lowercase: Package level variable

// Notes:
// No private scope, block level scope
// pascal or camelCase

func main() {
	fmt.Printf("%v, %T", I, I) // OP: 10, int
	fmt.Printf("%v, %T", i, i) // OP: 20, int

}
