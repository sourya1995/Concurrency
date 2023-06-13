package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func generateData() <-chan int64 {
	c := make(chan int64)
	const filePath = "integer.txt"
	go func() {
        file, _ := os.Open(filePath)
		defer close(c)
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")
			integer, _ := strconv.ParseInt(line, 10, 0)
			c <- integer
			if err == io.EOF {
                break
            }
		}
	}()

	return c
}