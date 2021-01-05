package main

import "fmt"

//We can return two objects in one function ---owSOME
// func test(x int, y int) (int, int) {
// 	return (x + y), (x - y)
// }

//It can also be written as
func test(x, y int) (z1, z2 int) {
	defer fmt.Println("hello") //It means that it will execute the statement after the return keyword inside that function is reached
	z1 = x + y
	z2 = x - y
	return
}
func main() {
	// -------------- one value -> func -> one value out
	ans1, ans2 := test(5, 9)
	fmt.Println(ans1)
	fmt.Println(ans2)

}
