package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Phone   string
	Email   string
	Address []Location
}

type Location struct {
	Street string
	City   string
	State  string
}

const (
	defaultCityIndex = 0
	ageOfAdulthood   = 18
)

func main() {
	person := Person{Name: "John Doe", Age: 25, Phone: "123-456-7890", Email: "john.doe@example.com", Address: []Location{{Street: "123 Main St", City: "Anytown", State: "NY"}}}
	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
	person.CelebrateBirthday()
}

func (p *Person) PrintPersonDetails() {
	fmt.Printf("Name: %s, Age: %d, Phone: %s, Email: %s, City: %s\n", p.Name, p.Age, p.Phone, p.Email, p.Address[defaultCityIndex].City)
}

func (p *Person) CelebrateBirthday() {
	p.Age += 1
	fmt.Println(p.formatAgeMessage())
}

func (p *Person) formatAgeMessage() string {
	return fmt.Sprintf("Happy Birthday %s! You are now %d years old.", p.Name, p.Age)
}

func (p *Person) IsAdult() bool {
	return p.Age >= ageOfAdulthood
}
