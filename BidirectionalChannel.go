package main

import "fmt"

func printValue(goChannel chan int){
	fmt.Println("Value inside the channel is", <-goChannel)
	goChannel <- 1

}

func main() {
	goChannel := make(chan int)
		go printValue(goChannel)
		goChannel <- 2
		fmt.Println("Value inside the channel is", <-goChannel)
}
