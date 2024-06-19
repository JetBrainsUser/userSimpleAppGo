package main

import (
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	Addresses []Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func (p *Person) DisplayDetails() {
	fmt.Printf("Name: %s, Age: %d, City: %s\n", p.Name, p.Age, p.Addresses[1])
}

func (p *Person) CelebrateBirthday() {
	p.Age += 1
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", p.Name, p.Age)
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	person := Person{Name: "John Doe", Age: 25, Addresses: []Address{}}

	person.DisplayDetails()

	person.CelebrateBirthday()

	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
}
