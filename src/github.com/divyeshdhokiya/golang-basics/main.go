package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	/* Numerated const */
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
