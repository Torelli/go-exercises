package main

import (
	"fmt"
)

func main() {
	x := 10
	y := "Hello!"
	z := 10.0

	fmt.Printf("x: %v, %T\n\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	x = x + int(z)
	fmt.Printf("x + z: %v, %T", x, x)
}
