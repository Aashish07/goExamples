//methods are some functions that we perform on some objects. Here we can make our own methods

package main

import "fmt"

type StudentMarks struct {
	name   string
	grades []int
	age    int
}

/*func (s1 StudentMarks) getAge() int {
	return s1.age
}

func (s1 StudentMarks) setAge(age int) {
	s1.age = age
}*/

func (s1 StudentMarks) averageGrade() float32 {
	sum := 0
	for _, val := range s1.grades {
		sum += val
	}
	return float32(sum) / float32(len(s1.grades))
}

func main() {
	s1 := StudentMarks{"Aashish", []int{99, 79, 99, 45, 45, 65, 67, 89, 23, 45}, 25}
	fmt.Println(s1)
	// s1.setAge(5)
	// fmt.Println(s1)
	average := s1.averageGrade()
	fmt.Println(average)
}
