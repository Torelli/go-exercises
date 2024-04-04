package main

import (
	"fmt"
)

type customType int

var x customType = 10

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Print(x)
}
