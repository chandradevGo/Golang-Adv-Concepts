package main

import "fmt"

// Define the 'person' struct
type person struct {
	Name string
}

// Define the 'secretAgent' struct, which embeds 'person'
type secretAgent struct {
	person
	ltk bool
}

// Implement the 'speak()' method for 'person'
func (p person) speak() {
	fmt.Println("I Need help -", p.Name)
}

// Implement the 'speak()' method for 'secretAgent'
func (sa secretAgent) speak() {
	fmt.Println("Don't worry, our secret agent", sa.Name, "will be there soon")
}

// Define the 'human' interface, which requires the 'speak()' method
type human interface {
	speak()
}

// Define a function that takes a 'human' interface and calls the 'speak()' method
func saysomething(h human) {
	h.speak()
}

func main() {
	// Create instances of 'person' and 'secretAgent'
	p1 := person{Name: "John"}
	sa := secretAgent{
		person: person{Name: "James Bond"},
		ltk:    true,
	}

	// Call 'saysomething' with both instances
	saysomething(p1)
	saysomething(sa)
}
