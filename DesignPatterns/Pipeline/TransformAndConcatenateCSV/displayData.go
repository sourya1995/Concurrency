package main

import "fmt"

func displayData(in chan []string) chan string {
	output := make(chan string)
	go func(in chan []string) {
		defer close(output)
        for c := range in {
			output <- fmt.Sprintf("First: %s, Second: %s, Third: %s", c[0], c[1], c[2])
		}
	}(in)

	return output
}