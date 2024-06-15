package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43
	// below one will be deadlock as we only init two buffer channel
	c <- 44

	fmt.Println(<-c)
	fmt.Println(<-c)
}
