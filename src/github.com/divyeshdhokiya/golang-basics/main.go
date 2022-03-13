package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* Convert int to float32 */
	var i int = 20
	fmt.Printf("%v, %T", i, i) // OP: 20, int

	var j float32 = 21
	/* Convert to float32 */
	j = float32(i)
	fmt.Printf("%v, %T", j, j) // OP: 20, float32

	/* Convert int to string */
	var l int = 22
	fmt.Printf("%v, %T", l, l) // OP: 20, int

	var k string = "str"
	k = strconv.Itoa(l)
	fmt.Printf("%v, %T", k, k) // OP: 22, string
}
