package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printValue(goChannel chan<-int) {
	goChannel <- 1
}

func main() {
	goChannel := make(chan int)

    go printValue(goChannel)
	fmt.Println("Value inside the channel is:", <-goChannel)
}
