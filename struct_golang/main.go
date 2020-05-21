package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//nikhil := person{firstName: "Nikhil", lastName: "Singh"}
	//fmt.Println(nikhil)

	// struct inside a struct
	nikhil := person{
		firstName: "Nikhil",
		lastName:  "Singh",
		contact: contactInfo{
			email:   "nikhilsingh.ns2014@gmail.com",
			zipcode: 263139,
		},
	}
	//fmt.Printf("%+v", nikhil)
	//pointer with struct
	//nikhilPointer := &nikhil  it works even if you don't define it and use nikhil directly
	nikhil.updateFirstName("Navyaa")
	nikhil.print()
}

//struct function with a pointer reciever
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// reciever with structs
func (p person) print() {
	fmt.Printf("%+v", p)
}
