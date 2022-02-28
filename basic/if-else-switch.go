package main

import "fmt"

func main() {

	num := 1
	// if num <= 5 {
	// 	fmt.Println("num is less than ", num)
	// } else {
	// 	fmt.Println("num is greater than", num)
	// }
	// if num == 1 {
	// 	fmt.Println("one")
	// } else if num == 2 {
	// 	fmt.Println("two")
	// } else {
	// 	fmt.Println("None")
	// }

	// this is the second method

	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("none")
	}

}
