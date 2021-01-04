package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6.5
	val := float32(x) != float32(y)
	fmt.Printf("%t", val)
	fmt.Println()
}

// different operators :- <,>,<=,>=,==,!=
