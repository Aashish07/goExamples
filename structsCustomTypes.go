//if you want your struct to be public, name it with caps
//if you want your struct to be private, name it small letters

package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func main() {
	//var x float32 = 5 //they have their own custom behaviour
	//we can make our own structure
	//using struct

	/*var p1 Point = Point{1, 2}
	var p2 Point = Point{-5, 7}
	p1.y = 10
	fmt.Println(p1.y)
	fmt.Println(p2.y)*/
	//p1 := Point{x: 3}

	//-------change value inside structure using reference pointer 23 to 27
	/*	p1 := &Point{x: 10, y: 20}
		p1.x = 20
		fmt.Println(*p1)
		change(p1)
		fmt.Println(*p1) */

	//----Struct inside a struct

	//p1 := &Point{y: 3}
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(*c1.centre)

}

type Circle struct {
	radius float64
	centre *Point
}

// func change(p1 *Point) {
// 	p1.x = 200
// }
