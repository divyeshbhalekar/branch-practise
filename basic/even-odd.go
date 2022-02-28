package main

import (
	"fmt"
)

func main() {
	num := 124 //(Creation & assignment on same line)
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

}
