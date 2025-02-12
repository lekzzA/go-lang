package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func main() {
	tanuj := person{
		firstName: "Tanuj",
		lastName:  "Bhatt",
		contactInfo: contactInfo{
			email:   "tanuj@.com",
			zipCode: 226022,
		},
	}
	tanuj.print()
	tanuj.updateName("Tannu")
	tanuj.print()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}
