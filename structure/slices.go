package main

import "fmt"

func main() {

	myslice := []string{"Hi",
		"Divyesh",
		"How",
		"are",
		"you"}

	updateslice(myslice)
	fmt.Println(myslice)

}

func updateslice(s []string) {
	s[0] = "hello"

}
