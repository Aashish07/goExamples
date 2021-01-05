//Range is the keyword in the language that allows us to actually
//iterate over items and elements inside of things like slices, arrays and strings

package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 5, 6, 7, 4, 3}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	//-------- another way of iterating   --i is represented as the index of the array
	// for i, element := range a {
	// 	fmt.Printf("%d: %d", i, element)
	// 	fmt.Println()
	// }
	//---------using only element variable, not using i
	for _, element := range a {
		fmt.Printf("%d", element)
		fmt.Println()
	}

}
