//Mutable(Changeable) and Immutable(not changable) Data Types

package main

import "fmt"

//func main() {
// 	var x []int = []int{3, 4, 5}
// 	y := x
// 	y[0] = 100
// 	fmt.Println(x, y)

// 	//slice is a mutable datatype, so value of x 'also' changed here

//--same goes with maps as well
// var x map[string]int = map[string]int{"Hello": 3}
// y := x
// y["y"] = 100
// x["7"] = 7
// fmt.Println(x, y)

//--Here also, any change in the value of either x or y will result in change of both of them

//They does that because the pointer in map and slice still points to the same old object

//and in arrays in doesNT

//--let us focus on arrays
//arrays store duplicate objects

/*var x [2]int = [2]int{3, 4}

//and if i write   x := []int{3,4} --> this is a slice

y := x

y[0] = 100

fmt.Println(x, y)*/

func changeFirst(slice []int) {
	slice[0] = 1000
}
func main() {
	var x []int = []int{3, 4, 5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}

//}
