package main

import "fmt"

func main() {
	/* Text */
	s := "This is string"
	fmt.Printf("%v,%T\n", s, s)                       // OP: This is string,string
	fmt.Printf("%v,%T\n", string(s[1]), string(s[1])) // OP: h,string

	// Concat
	s1 := "str1"
	s2 := "str2"
	fmt.Printf("%v,%T", s1+s2, s1+s2) // OP: str1str2,string

	// Convert to collection of bytes
	bs1 := "str1"
	bs2 := []byte(bs1)
	fmt.Printf("%v,%T", bs2, bs2) // OP: [115 116 114 49],[]uint8

}
