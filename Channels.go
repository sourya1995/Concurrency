package main

import "fmt"

func main() {
	goChannel := make(chan int)
	go func() {
		goChannel <- 1
	}()//concurrent goroutine sends value to channel

	
	fmt.Println("Value", <-goChannel)
}