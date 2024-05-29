package main

import "fmt"

type person struct {
	Name string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I Need help -", p.Name)
}

func (sa secretAgent) speak() {
	fmt.Println("Don't worry secret Agent -", sa.Name, "will be there soon")
}

type human interface {
	speak()
}

func command(h human) {
	h.speak()
}

func main() {
	p1 := person{
		Name: "Sam",
	}

	sa := secretAgent{
		person: person{
			Name: "James",
		},
		ltk: true,
	}

	command(p1)
	command(sa)

}
