//Arrays are ordered collection of elements

package main

import "fmt"

func main() {
	//We need to make sure that everyting we store in an array is of the same type ( var arr []String)
	// var arr [5]int //array is defined ----------------------
	// arr[2] = 7
	// fmt.Println(arr[2])

	//another way of defining array --------------
	// arr := [3]int{3, 4, 5}
	// sum := 0
	// for i := 0; i < (len(arr)); i++ {
	// 	sum += arr[i]
	// }
	// fmt.Println(sum)

	//2d array

	arr2D := [2][3]int{{1, 2, 4}, {3, 4, 6}}
	fmt.Println(arr2D[0][2])

}
