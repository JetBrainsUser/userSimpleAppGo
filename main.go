package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address []string
}

type Location struct {
	Street string
	City   string
	State  string
}

const detailsFormat = "Name: %s, Age: %d, City: %s\n"

func (p *Person) displayCity() string {
	return p.Address[1]
}

func (p *Person) DisplayDetails() {
	city := p.displayCity()
	fmt.Printf(detailsFormat, p.Name, p.Age, city)
}

func (p *Person) CelebrateBirthday() {
	p.Age += 1
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", p.Name, p.Age)
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	person := Person{Name: "John Doe", Age: 25, Address: []string{"123 Main St", "Anytown", "NY"}}

	person.DisplayDetails()

	person.CelebrateBirthday()

	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
}
