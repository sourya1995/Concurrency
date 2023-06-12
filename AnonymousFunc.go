package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() {

	g := func(v string) string {
		fmt.Println(v + "Dear!!")
		return v
	}

	anonymous := func(v string, g func(v string)string) {
		x := g(v)
		fmt.Println(x)
		wg.Done()
	}

	wg.Add(1)
	go anonymous("hello", g)
	wg.Wait()
}