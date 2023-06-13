package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	jobs := make(chan int, 10)
	done := make(chan bool)

	go func(){
		for {
			j, more := <-jobs
			if more{
				fmt.Println(j)
			} else {
				done <- true
				fmt.Println("Received all jobs")
				return
			}
		}
	}()

	wg1.Add(1)
	wg2.Add(1)
	go func(){
		defer wg1.Done()
        for i := 0; i <= 100; i++ {
            jobs <- i
			fmt.Println("sending job", i)
        }
	}()

	go func(){
		defer wg2.Done()
        wg1.Wait()
		close(jobs)
        fmt.Println("Sent all jobs")
        <-done
        }()
		wg2.Wait()
	}

