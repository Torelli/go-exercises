package main

import (
	"fmt"
)

func main() {
	x := 200
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %#x\n", x)
	y := x << 1
	fmt.Printf("Decimal: %d\n", y)
	fmt.Printf("Binary: %b\n", y)
	fmt.Printf("Hex: %#x\n", y)
	z := x >> 1
	fmt.Printf("Decimal: %d\n", z)
	fmt.Printf("Binary: %b\n", z)
	fmt.Printf("Hex: %#x\n", z)
}
