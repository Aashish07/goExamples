package main

import "fmt"

func main() {
	//defer function in case there is some eof execution we want, in any scenario
	//panic - It is mostly like exception
	//recover -> It is effective only inside defer function
	//Once inside a function panic happened, the request for the function stops and goes to the defer function and then recover can handle it

	defer Welcome("2020")
	panic("Weird Error")
}

func Welcome(str string) {
	fmt.Println("Past year was " + str)
	if r := recover(); r != nil {
		fmt.Println("2020 Recovered from corona")
	}
}
