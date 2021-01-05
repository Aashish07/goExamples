package main

import "fmt"

/*
	& stands for get the pointer   ->
	&x gives us memory location i.e reference to x
	* represents for derefrence --> return value at location i.e value at &x(address)

*/

// func main() {
// 	x := 7
// 	y := &x
// 	fmt.Println(x, y)
// 	*y = 8 //means *(&x) = *y
// 	fmt.Println(x, y)
// }

//This example change a value of variable by using its reference value
/*func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)
}*/
func main() {
	toChange := "hello"
	var pointer *string = &toChange
	fmt.Println(*pointer)
}
