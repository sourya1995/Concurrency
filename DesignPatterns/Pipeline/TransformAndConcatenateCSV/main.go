package main

import (
  "fmt"
  "log"
)

func main(){
  //c := displayData(convertTwo(convertOne(generateData("file.csv"))))
  records, err := generateData("file.csv")
  if err != nil {
    log.Fatalf("Could not read csv %v", err)
  }
  for data := range displayData(convertTwo(convertOne(records))) {
    fmt.Println(data)
  }
}