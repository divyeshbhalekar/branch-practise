package main

import "fmt"

// var a is defined twice one inside the funct & one outside
//the var inside func is always take prefernce.
//if there is no variable inside the fun it will look for variable in package
//

var a = 8 // package level variable it is available throughout the code

func demo() {
	var a = 8 // function level variable it is available only inside the function
	fmt.Println(a)
}

func main() {
	demo()
	fmt.Println(a)

}
