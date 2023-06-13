package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Number struct {
	square, wId int
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {

		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func squareNumber(in <- chan int, val int) <- chan Number{
	out := make(chan Number)
    go func(val int) {
        for n := range in {
            out <- Number{n * n, val,}
        }
        close(out)
    }(val)
    return out
}

func displayData(cs ...<-chan Number) {
	for _, val := range cs {
		wg.Add(1)
		go func(val <-chan Number) {
			defer wg.Done()
            for val := range val {
                fmt.Println(val.square, val.wId)
            }
        }(val)
	}
}

func main() {
	randomNumbers := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan := generatePipeline(randomNumbers)

	// Fan-out to 3 Go-routine
	c1 := squareNumber(inputChan, 1)
	c2 := squareNumber(inputChan, 2)
	c3 := squareNumber(inputChan, 3)

	displayData(c1, c2, c3)
	wg.Wait()
}
