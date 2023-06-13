package main

import (
	"log"
)

func main() {

	c := displayData(prepareData(generateData()))
	for data := range c {
		log.Printf("Items: %+v", data)
	}

}