package main

import "fmt"

func main() {

	// defer help to run skip and execute at the end .
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello there")
}
