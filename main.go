package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address []Location
}

type Location struct {
	Street string
	City   string
	State  string
}

const (
	cityIndex      = 1
	ageOfAdulthood = 18
)

func main() {
	person := Person{Name: "John Doe", Age: 25, Address: []Location{{Street: "123 Main St", City: "Anytown", State: "NY"}}}

	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
	person.CelebrateBirthday(person.Name)

}

func (person *Person) PrintPersonDetails() {
	cityName := person.Address[cityIndex].City
	fmt.Printf("Name: %s, Age: %d, City: %s\n", person.Name, person.Age, cityName)
}

func (p *Person) CelebrateBirthday(name string) {
	p.Age += 1
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", name, p.Age)
}

func (p *Person) IsAdult() bool {
	return p.Age >= ageOfAdulthood
}
