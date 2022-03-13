package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

type Car struct {
	Model string
	Year  int
	Color string
	Parts []string
}

func main() {
	/* Struct */
	aCar := Car{
		Model: "Toyota",
		Year:  2000,
		Color: "White",
		Parts: []string{
			"Doors",
			"Wheels",
		},
	}

	fmt.Println(aCar)

	// Fetch
	fmt.Println(aCar.Model)    // OP: Toyota
	fmt.Println(aCar.Parts)    // OP: [Doors Wheels]
	fmt.Println(aCar.Parts[1]) // OP: Wheels

	/* Anonymous struct */
	aDr := struct{ name string }{name: "Dr X"}
	fmt.Println(aDr) // OP: {Dr X}
	bDr := aDr
	bDr.name = "Dr Z"
	fmt.Println(bDr) // OP: {Dr Z}

	// Note: use & to manipulate main struct

	/* Embedded */
	b := Bird{}
	b.Name = "Sparrow"
	b.Origin = "IN"
	b.CanFly = true
	b.Speed = 10

	fmt.Println(b) // OP: {{Sparrow IN} 10 true}

}
