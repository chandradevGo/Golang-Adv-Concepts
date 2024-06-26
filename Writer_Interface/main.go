package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// this is to write to buffer
type person struct {
	Name string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.Name))
	return err
}

// using below method we can create two text files and write some data to it.
// capture any erros encountered
func main() {
	// creat and write to files f1 & f2
	f1, err := os.Create("hello.text")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f1.Close()

	s := []byte("Hello gophers!")

	_, err = f1.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	f2, err := os.Create("success.text")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f2.Close()

	s2 := []byte("You have succesfully created two text files and wrote some data to those.")

	_, err = f2.Write(s2)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// writes to buffer and prints buffer from in-memory
	p := person{
		Name: "James",
	}

	var b bytes.Buffer
	p.writeOut(&b)
	fmt.Println(b.String())
}
