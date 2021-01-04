package main

import (
	"fmt"
)

func main() {
	name := "Aashish"
	fmt.Println("Before if")
	if name != "Aashish" {
		fmt.Println("Inside if")
	}
	fmt.Println("After if")
}
