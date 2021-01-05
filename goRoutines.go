package main

import (
	"fmt"
	"time"
)

// A goRoutine is a lightweight thread managed by the go runtime
//goRoutines runs in the same address space

//concurrency in the program will be received by writing the go routines

//---Normal Example -> less efficient way
/*func main() {
	fmt.Println("Hello")

	func() {
		fmt.Println("Anonymous")
		time.Sleep(4 * time.Second)
	}()

	namedFunction()
	fmt.Println("Final Statement.....")

}*/

func namedFunction() {
	fmt.Println("Call 02 from named function")
	time.Sleep(4 * time.Second)
}

func main() {
	fmt.Println("Hello")

	go func() { //this go keyword will help program run asyncronously
		fmt.Println("Call 01 from Anonymous function")
		time.Sleep(3 * time.Second)
	}()

	go namedFunction()
	time.Sleep(10 * time.Second)
	fmt.Println("Final Statement.....")
}
