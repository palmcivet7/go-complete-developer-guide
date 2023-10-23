package main

import "fmt"

type contactInfo struct {
	email	string
	zipCode	int
}

type person struct {
	firstName	string
	lastName	string
	// contact		contactInfo
	contactInfo
}

func main() {
	// palmcivet := person{firstName: "Palm", lastName: "Civet"}
	// fmt.Println(palmcivet)

	// var palmcivet person
	// palmcivet.firstName = "Palm"
	// palmcivet.lastName = "Civet"

	// fmt.Println(palmcivet)
	// fmt.Printf("%+v", palmcivet)

	palmcivet := person{
		firstName: "Palm",
		lastName: "Civet",
		contactInfo: contactInfo{
			email: "palmcivet@proton.me",
			zipCode: 12345,
		},
	}

	palmcivet.updateName("Palmy")
	palmcivet.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName 
}

func (p person) print() {
	fmt.Printf("%+v", p)
}