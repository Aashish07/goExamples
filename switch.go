package main

import (
	"fmt"
)

func main() {
	ans := -1

	// switch ans {
	// case 1, -1: //advantage with this one
	// 	fmt.Println("1. one")
	// 	fmt.Println("2. one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not a case")
	// }

	switch {
	case ans < 0:
		fmt.Println("lesser than 0")
	}
}
