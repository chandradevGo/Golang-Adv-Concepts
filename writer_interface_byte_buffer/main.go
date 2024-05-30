package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("Gophers!")
	fmt.Println(b.String())

	b.Reset()

	b.WriteString("It's Thursday ")
	fmt.Println(b.String())

	b.WriteString("Happy Happy")
	fmt.Println(b.String())
}
