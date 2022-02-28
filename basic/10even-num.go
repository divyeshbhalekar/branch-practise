package main

import "fmt"

func main() {

	var counter int = 0

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			if counter == 10 {
				break
			} else {
				counter++
			}
			fmt.Println(i, "is even")
		} /*else {
			fmt.Println(i, "is odd")
		}*/

	}

	/*fmt.Println("Enter number: ")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0; n <= 5;n++ {
		fmt.Println(n, " is even number")
	} else {
		fmt.Println(n, "is odd number")

	}*/

}
