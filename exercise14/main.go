package main

import (
	"fmt"
)

const (
	_ = 2024 + iota
	next1
	next2
	next3
	next4
)

func main() {
	fmt.Println(next1, next2, next3, next4)
}
