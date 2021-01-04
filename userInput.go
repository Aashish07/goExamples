package main

import (
	"bufio" //buffered input/output
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born : ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //_ defined in the place of error if it get raised from the statement
	fmt.Printf("You will be %d at the end of 2020", 2020-input)
	fmt.Println()
}
