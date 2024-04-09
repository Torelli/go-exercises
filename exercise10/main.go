package main

import (
	"fmt"
)

var x = 10
var y = 4

func main() {
	a := x == y
	b := x != y
	c := x <= y
	d := x < y
	e := x >= y
	f := x > y

	fmt.Printf("a: %v\n b: %v\n c: %v\nd: %v\ne: %v\nf: %v", a, b, c, d, e, f)
}
