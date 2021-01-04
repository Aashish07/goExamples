package main

import (
	"fmt"
)

func main() {
	// x := 0
	// //for loop --------------------1
	// for x <= 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	//another for loop -----------------2
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			break
			//break will make print only 1(first)
			//continue will make print all of them till 1000
		}
		//fmt.Println("N")
	}

	//break immediately exits the loop skips the eol bracket
	// continue goes back to beginning of loop, so anything below it does not happen
}
