package main

import (
	"fmt"
	"math"
)

func main() {

	var num float64 = 12

	var result = math.Sqrt(num)
	var intresult = math.Round(result)

	fmt.Println("hello")
	// if we write %f it will give you rounded up value of 3.5
	fmt.Printf("the o/p is %.2f thankyou", result)
	println()
	// this gives the o/p in round off
	fmt.Printf("the op is %g ", intresult)
}
