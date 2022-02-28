package main

import "fmt"

func calc(x, y int) (int, int) {
	out1 := x + y
	out2 := x - y
	return out1, out2
}

func main() {

	var a = 3
	var b = 2
	c, d := calc(a, b)
	fmt.Println(c, "\n", d)
}
