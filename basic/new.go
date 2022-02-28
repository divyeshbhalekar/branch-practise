// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello world")
// }

package main

import "fmt"

func main() {
	fmt.Println(5 + 4)

	// var a = 25
	// var b = 26
	// var c = (a + b)
	// // fmt.Print(c)
	// println(c)

	var a, b int
	a, b = 5, 6
	var c = a + b
	println(c)

	for i := 1; i <= 10; i++ {
		println("hello", i)
	}
}
