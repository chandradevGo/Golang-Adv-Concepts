package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
	}

	p2 := person{
		FirstName: "Miss",
		LastName:  "MonneyPenny",
		Age:       27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	pp, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(pp))

}
