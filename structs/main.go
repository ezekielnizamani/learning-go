package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	hizqeel := person{
		firstName: "hizqeel",
		lastName:  "ahmed ali",
		contact: contactInfo{
			email:   "ezekielnizamani0@gmail.com",
			zipCode: 700111,
		},
	}
	hizqeel.updateName("adam")
	hizqeel.print()

}
func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
