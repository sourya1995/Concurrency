package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	val int
	message string
}

func main() {
	for returnedVal := range randomGenerator(10) {
		fmt.Println(returnedVal )
	}
}
func randomGenerator(n int) <-chan node {
	output := make(chan node)
	go func() {
		defer close(output)
		fmt.Println("I am the producer")
		for i = 0; i < n; i++ {
			num:= rand.Intn(n) + 1
			output <- random {
				val: num,
				message: fmt.Sprintf("Hello %d", num),
			}
		}
	}()
	return output

		
	
	
}

