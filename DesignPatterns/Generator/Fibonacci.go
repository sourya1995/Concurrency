package main

import (
	"fmt"
	"sync"

)

var wg sync.WaitGroup
func main() {
	for c := range(1000000){
		fmt.Printf("Fibonacci number: %d\n", c)
	}
}


func fibonacci(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		fmt.Println("hello, this is the producer")
		for i, j := 0, 1; i < n; i, j = i+j, i{
			out <- i
		}
	}()
	return out
}