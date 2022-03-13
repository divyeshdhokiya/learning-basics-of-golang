package main

import "fmt"

/* Package level variables. can not use ":=" syntax at this place */
var (
	p   int    = 30
	str string = "Hi!"
)

func main() {
	var num int
	/* Print default value */
	fmt.Println(num) // OP: 0

	/* Ways of declaring variables */
	/* 1: Need to declare but not ready to initialize it right away */
	var i int
	i = 20
	// Examples
	fmt.Println(i) // OP: 20

	/* 2: When go does not have enough information of the type of the value assigned to the variable */
	var j int = 22
	var l float32 = 1.5
	// Examples
	fmt.Printf("%v, %T", j, j) // OP: 22, int
	fmt.Printf("%v, %T", l, l) // OP: 1.5 float32

	/* 3: short and clean way */
	k := 21
	fmt.Printf("%v, %T", k, k)

	fmt.Printf("%v, %T", p, p)
	fmt.Printf("%v, %T", str, str)

}
