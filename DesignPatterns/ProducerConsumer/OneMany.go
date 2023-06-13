package main

import (
  "fmt"
  "time"
  "sync"
  "math/rand"
)

func increment(previous int) int {
	time.Sleep(time.Duration((rand.Intn(100) + 1)) * time.Millisecond)
	return previous + 1
}

func main() {
	noOfConsumer := 10
	var wg sync.WaitGroup
	data := make(chan int)

	go func() {
		defer close(data)
		for i := 0; i < 100; i++ {
			data <- increment(i)
        }
	}()


	for i := 0; i < noOfConsumer; i++ {
		wg.Add(i)
		go func(i int) {
			defer wg.Done()
            for data := range data {
                fmt.Println(data, i)
            }
        }(i)
	}
	wg.Wait()
}