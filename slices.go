//slices take a portion of the array
//we can change the size of the slices  --pointer,length,capacity
package main

import "fmt"

func main() {

	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3] //if i dont put a size on lhs inside [], it means it is a slice.
	// //if we dont have anything on rhs, it means just copy [1:3] means start at 1 go to 3.
	// fmt.Println(s[:cap(s)])

	// var a []int = []int{5, 6, 7, 8, 9}
	// b := append(a, 10) //an advantage over arrays, adding an extra adding int
	// fmt.Println(b)

	//use slices with make
	a := make([]int, 5)
	fmt.Printf("%T", a)
	fmt.Println()

}
