package main

import "fmt"

func main() {
	/* Slices */
	students := []int{1, 2, 3}
	fmt.Printf("%v\n", students)      // OP: [1 2 3]
	fmt.Printf("%v\n", len(students)) // OP: 3
	fmt.Printf("%v\n", cap(students)) // OP: 3

	a := []int{11, 22, 33, 44, 55, 66, 77, 88}
	b := a[:]
	c := a[:3]
	d := a[5:]
	e := a[5:7]
	fmt.Printf("%v\n", a) // OP: [11 22 33 44 55 66 77 88]
	fmt.Printf("%v\n", b) // OP: [11 22 33 44 55 66 77 88]
	fmt.Printf("%v\n", c) // OP: [11 22 33]
	fmt.Printf("%v\n", d) // OP: [66 77 88]
	fmt.Printf("%v\n", e) // OP: [66 77]

	a[5] = 100            // Update original array hence effect to all other arrays
	fmt.Printf("%v\n", a) // OP: [11 22 33 44 55 100 77 88]
	fmt.Printf("%v\n", b) // OP: [11 22 33 44 55 100 77 88]
	fmt.Printf("%v\n", c) // OP: [11 22 33]
	fmt.Printf("%v\n", d) // OP: [100 77 88]
	fmt.Printf("%v\n", e) // OP: [100 77]

	// Make function
	aa := make([]int, 3, 100)
	fmt.Printf("%v\n", aa)      // OP: [0 0 0]
	fmt.Printf("%v\n", len(aa)) // OP: 3
	fmt.Printf("%v\n", cap(aa)) // OP: 100

	// ---
	x := []int{}
	fmt.Printf("%v\n", x)      // OP: []
	fmt.Printf("%v\n", len(x)) // OP: 0
	fmt.Printf("%v\n", cap(x)) // OP: 0

	x = append(x, 2)
	fmt.Printf("%v\n", x)      // OP: [2]
	fmt.Printf("%v\n", len(x)) // OP: 1
	fmt.Printf("%v\n", cap(x)) // OP: 1

	// x = append(x, 3, 4, 5, 6)
	// fmt.Printf("%v\n", x)      // OP: [2 3 4 5 6]
	// fmt.Printf("%v\n", len(x)) // OP: 5
	// fmt.Printf("%v\n", cap(x)) // OP: 6

	x = append(x, []int{3, 4, 5, 6}...)
	fmt.Printf("%v\n", x)      // OP: [2 3 4 5 6]
	fmt.Printf("%v\n", len(x)) // OP: 5
	fmt.Printf("%v\n", cap(x)) // OP: 6

	// Stack
	w := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", w) // OP: [1 2 3 4 5 6]
	// Shift: remove first element
	y := w[1:]
	fmt.Printf("%v\n", y) // OP: [2 3 4 5 6]
	// Shift: remove last element
	y2 := w[:len(w)-1]
	fmt.Printf("%v\n", y2) // OP: [1 2 3 4 5]
	// Shift: remove n element -> 3
	y3 := append(w[:2], w[3:]...)
	fmt.Printf("%v\n", y3) // OP: [1 2 4 5 6]

	
}
