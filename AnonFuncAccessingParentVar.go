package main

import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i:=1; i<=10; i++ {
		wg.Add(1)
		go func(i int){
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}