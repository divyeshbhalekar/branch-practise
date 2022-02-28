package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   int
}

type contactinfo struct {
	email   string
	zipcode int
}

func main() {

	var divyesh = []person{
		{firstname: "divyesh", lastname: "bhalekar", contact: 842},
	}

	updatename(divyesh)
	fmt.Println(divyesh)
}
func updatename(s person) {
	s[0] = "Divyesh"

}
