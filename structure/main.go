package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactinfo
}

type contactinfo struct {
	email   string
	zipcode int
}

func main() {

	divyesh := person{
		firstname: "divyesh",
		lastname:  "bhalekar",
		contact: contactinfo{
			email:   "divyesh@gmail.com",
			zipcode: 400084,
		},
	}

	divyesh.updatename("Divyesh", "Bhalekar")
	divyesh.print()

}
func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p *person) updatename(newfirstname, newlastname string) {
	(*p).firstname = newfirstname
	(*p).lastname = newlastname
}
