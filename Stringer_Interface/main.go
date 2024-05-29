package main

import (
	"fmt"
	"strconv"
)

// Define a struct named 'book' with a single field 'title'
type book struct {
	title string
}

// Implement the String method for the book struct to satisfy the fmt.Stringer interface
func (b book) String() string {
	return fmt.Sprint("The title of the book is: ", b.title)
}

// Define a new type 'count' based on the built-in int type
type count int

// Implement the String method for the count type to satisfy the fmt.Stringer interface
func (c count) String() string {
	return fmt.Sprint("The count is: ", strconv.Itoa(int(c)))
}

// Define a function logInfo that takes an argument of type fmt.Stringer
// This function will print a log message including the string representation of the argument
func logInfo(s fmt.Stringer) {
	fmt.Println("Log from stringer wrapper", s.String())
}

func main() {
	// Create an instance of the book struct with a specific title
	b := book{
		title: "Cracking the coding interview",
	}

	// Call logInfo with the book instance, which will use the book's String method
	logInfo(b)

	// Create a variable of type count with a value of 5
	var c count = 5

	// Call logInfo with the count instance, which will use the count's String method
	logInfo(c)
}
