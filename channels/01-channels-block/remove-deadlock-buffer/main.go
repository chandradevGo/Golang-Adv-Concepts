package main

import "fmt"

func processChannel() []int {
	c := make(chan int, 2)

	c <- 42
	c <- 43
	// The following line would cause a deadlock if uncommented
	// c <- 44

	result := []int{<-c, <-c}
	return result
}

func main() {
	results := processChannel()
	for _, result := range results {
		fmt.Println(result)
	}
}
