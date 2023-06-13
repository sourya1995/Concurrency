package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q, a string
}

func check(e error) {
	if e!= nil {
        panic(e)
    }
}

func readCsvFile(filename string) [][]string {
	data, err := os.Open(filename)
	check(err)
	defer data.Close()

	csvReader := csv.NewReader(data)
	records, err := csvReader.ReadAll()
	check(err)
	return records
}

func parseProblems(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i := 0; i < len(records); i++ {
		problems[i] = problem{
            q: records[i][0],
            a: records[i][1],
        }
		return problems
}

func main(){
	records := readCsvFile("input.csv")
    problems := parseProblems(records)
	timer := time.NewTimer(time.Duration(5) * time.Second)
	correct, incorrect := 0, 0
	answerCh := make(chan string)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s ", i+1, p.q)
		go func() {
            var answer string
            fmt.Scanf("%s", &answer)
			answerCh <- answer
        }()

		select {
			case <- timer.C:
				fmt.Printf("\nCorrect %d, Incorrect %d, Total %d\n", correct, incorrect, len(problems))
				return
			case answer := <-answerCh:
				if answer == p.a {
					correct++
				} else {
					incorrect++
				}
			}
		}
    fmt.Printf("\nCorrect %d, Incorrect %d, Total %d\n", correct, incorrect, len(problems))
}