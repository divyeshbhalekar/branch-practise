package main

import "fmt"

type student struct {
	id   int
	name string
	city string
}

func main() {

	var a = student{1, "divyes", "mumbai"}
	//	var a student = student{1, "divyes", "mumbai"}   (this is same as above line)

	fmt.Println(a)
	fmt.Println(a.name) // to print only name

	var b = student{2, "sagar", "thane"}
	fmt.Println(b)
}
