package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { //this is a goroutine
			for c := 0; c < 100; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
		wg.Wait() //wait for all waitGroups to finish
		fmt.Println("ops: ", atomic.LoadUint64(&ops))
	}
