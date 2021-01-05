package main

import "fmt"

//functions have there own datatype

// func main() {
// 	ask := func(x int) int {
// 		return x * -1
// 	}(8) //----we can pass the value like that as well

// 	fmt.Println(ask)

// }

//Lets now pass a function to another function as a paramtere ---------
// func quesn(ask func(int) int) {
// 	fmt.Println(ask(8))
// }

// func main() {
// 	ask := func(x int) int {
// 		return x * -1
// 	}
// 	quesn(ask)
// }

//I can also call a function immediately

func main() {
	func() {
		fmt.Println("test")
	}()
}
