package main

import "fmt"

const (
	_ = iota + 5
	a
	b
	c
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

const (
	isAdmin = 1 << iota
	isUser
)

func main() {
	/* Numerated const */
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Example 1
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	fmt.Printf("%.2fMB\n", fileSize/MB)

	// Example 2
	var roles byte = isAdmin
	fmt.Printf("isAdmin: %v\n", isAdmin&roles == isAdmin) // true
	fmt.Printf("isUser: %v\n", isUser&roles == isUser)    // false
}
