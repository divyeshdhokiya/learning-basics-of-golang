package main

import "fmt"

func main() {
	/* Arrays */
	// 1. predefined size
	students := [3]int{1, 2, 3}
	fmt.Printf("%v\n", students) // OP: [1 2 3]

	// 2. dynamic size
	dStudents := [...]int{1, 2, 3, 4}
	fmt.Printf("%v\n", dStudents) // OP: [1 2 3 4]

	arr := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", arr) // OP: [1 2 3 4]

	var students1 [3]string
	fmt.Printf("%v\n", students1) // OP: [ ]
	students1[0] = "First"
	students1[1] = "Second"
	students1[2] = "Third"
	fmt.Printf("%v\n", students1)      // OP: [First Second Third]
	fmt.Printf("%v\n", len(students1)) // OP: 3

	// ---
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Printf("%v\n", matrix) // OP: [[1 2 3] [4 5 6] [0 0 0]]

	matrix[2] = [...]int{6, 7, 8}
	fmt.Printf("%v\n", matrix) // OP: [[1 2 3] [4 5 6] [6 7 8]]

	arr1 := [...]int{1, 2, 3}

	// Copy array: value
	arr2 := arr1
	arr2[2] = 4

	fmt.Printf("%v\n", arr1) // OP: [1 2 3]
	fmt.Printf("%v\n", arr2) // OP: [1 2 4]

	// Copy array: reference
	arr3 := &arr1
	arr3[1] = 9

	fmt.Printf("%v\n", arr1) // OP: [1 9 3]
	fmt.Printf("%v\n", arr3) // OP: &[1 9 3]

}
