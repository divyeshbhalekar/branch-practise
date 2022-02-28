package main

import "fmt"

func main() {

	// there are different ways to declare variable below we will see

	var num1 = 1     // 1 way
	var num2 int = 3 // 2 way

	var num3 int // 3 way (By default int value is 0)
	num3 = 5

	var a, b int //(Declaring two variables on the same line)
	a, b = 9, 1
	var c = a + b
	println(c)

	// to add variables
	var result int = (num1 + num2)

	// to declare a constant (a const cannot change)
	const num4 int = 6

	fmt.Println(num1, "\n", num2, num3, result, num4) //"\n" to get o/p to next line
	// var a = 1
	fmt.Println(a)
}
