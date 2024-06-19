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

func (p *Person) PrintDetails() {
	city := p.Address[1]
	fmt.Printf("Name: %s, Age: %d, City: %s\n", p.Name, p.Age, city)
}

const ageOfAdulthood = 18

func (p *Person) celebrateBirthday() {
	p.incrementAge()
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", p.Name, p.Age)
}

func (p *Person) incrementAge() {
	p.Age += 1
}

func (p *Person) IsAdult() bool {
	return p.Age >= ageOfAdulthood
}

func main() {
	person := Person{Name: "John Doe", Age: 25, Address: []string{"123 Main St", "Anytown", "NY"}}

	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
}
