package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Number struct {
	original, reverse int
}

func readFile(filename string) (<-chan int, error) {
	out := make(chan int)
	file, error := os.Open(filename)
	if error!= nil {
        return nil, error
    }

	go func(file *os.File) {
		defer close(out)
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSuffix("\n")
			line = strings.TrimSuffix("\r")
			integer, _ := strconv.Atoi(line)
			out <- integer
			if err == io.EOF {
                break
            }
		}
	}(file)
	return out, nil
}

func main() {
	// reading file
	channel1, err := readFile("file1.txt")
	if err != nil {
		panic(err)
	}

	channel2, err := readFile("file2.txt")
	if err != nil {
		panic(err)
	}

	channel3, err := readFile("file3.txt")
	if err != nil {
		panic(err)
	}

	channel := merge(channel1, channel2, channel3)

	for val := range channel {
		fmt.Printf("Original Number is %d, Reverse number is %d\n", val.original, val.reverse)
	}
}

func merge(cs ...<-chan int) <-chan Number {
	var wg sync.WaitGroup
	out := make(chan Number)

	send := func(c <- chan int){
		for n := range c {
			out <- Number{
				n, reverseNumber(n),
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
        wg.Wait()
        close(out)
    }()

	return out

}

func reverseNumber(n int) int {
	reverse := 0
	for n != 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n = n / 10
	}
	return reverse
}

