package main

import "fmt"

func main() {
	/* Panic:  how app can panic*/
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans) // OP: panic: runtime error: integer divide by zero

	/* 	fmt.Println("start")
	   	panic("something went wrong!")
	   	fmt.Println("end")

	   	/*
	   		OP:
	   			start
	   			panic: something went wrong!
	*/

	fmt.Println("start")
	defer fmt.Println("Deferred")
	panic("something went wrong!")
	fmt.Println("end")

	/*
		OP:
			start
			Deferred
			panic: something went wrong!
	*/

}
