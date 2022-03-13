package main

import "fmt"

func main() {
	/* Maps */
	/* 1 */
	statePopulations := map[string]int{
		"a": 10,
		"b": 10,
		"c": 10,
		"d": 10,
	}
	fmt.Println(statePopulations) // OP: map[a:10 b:10 c:10 d:10]

	/* 2 */
	statePops := make(map[string]int)
	statePops = map[string]int{
		"a": 10,
		"b": 10,
		"c": 10,
		"d": 10,
	}
	fmt.Println(statePops) // OP: map[a:10 b:10 c:10 d:10]

	// Read one
	fmt.Println(statePops["a"]) // OP: 10

	// Add new
	statePops["e"] = 10
	fmt.Println(statePops) // OP: map[a:10 b:10 c:10 d:10, e:10]

	//Delete
	delete(statePops, "e")
	fmt.Println(statePops) // OP: map[a:10 b:10 c:10 d:10]

	fmt.Println(statePops["e"]) // OP: 0

	// is exist
	key, ok := statePops["e"]
	fmt.Println(key, ok) // OP: 0, false

}
