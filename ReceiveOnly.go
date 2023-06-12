package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printValue(goChannel <- chan int) {
	fmt.Println("Value inside the channel is: ", <-goChannel)
	wg.Done()
} 

func main() {
	goChannel := make(chan int)
	wg.Add(1)
	go printValue(goChannel)
	goChannel <- 2
	wg.Wait()
}