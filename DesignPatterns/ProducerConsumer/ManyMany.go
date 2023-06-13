package main

import (
	"crypto/x509"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	noOfProducer := 10
	noOfConsumer := 10
	data := make(chan int64)
	var producerGroup sync.WaitGroup

	//Producer
	var ops int64
	for i := 0; i < noOfProducer; i++ {
		producerGroup.Add(1)
		go func() {
			defer producerGroup.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&ops, 1)

			}
		}()
	}

	go func() {
		defer close(data)
		producerGroup.wait()
	}()

	var consumerGroup sync.WaitGroup
	for i := 0; i < noOfConsumer; i++ {
		consumerGroup.Add(1)
		go func(i int) {
			defer consumerGroup.Done()
			for data := range data {
				fmt.Printf("Value of i = %d Printed by consumer %d\n", data, i)
			}
		}(i)
	}
	consumerGroup.Wait()
}
