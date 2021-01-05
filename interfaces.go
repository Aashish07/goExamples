//interfaces - way of looking at a set of related object and types

//This "interface will help us implement the behavioural change between different function with same name"
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func getArea(s shape) float64 {
	return s.area()
}
func main() {
	c1 := Circle{4.5}
	r1 := Rectangle{5, 7}
	shapes := []shape{c1, r1}
	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
