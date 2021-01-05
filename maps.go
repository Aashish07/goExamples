package main

import (
	"fmt"
)

func main() {
	// var mp map[string]int = map[string]int{
	// 	"apple":  5,
	// 	"pear":   6,
	// 	"orange": 7, //{9,5,3,5,5,23}
	// }

	//other way to make map
	mp := make(map[string]int)
	//mp["apple"] = 78
	mp["aashish"] = 32
	fmt.Println(mp)

	fmt.Printf("Length of map is: %d", len(mp))
	fmt.Println()

	// -- for deletetion --
	// delete(mp, "apple")
	// fmt.Println(mp)

	//--if a key exists --
	val, ok := mp["aashish"]
	fmt.Println(val, ok)
}
