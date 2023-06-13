package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	randomno int
}

type Result struct {
	job Job
	sumofdigits, wId int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
        sum += digit
        no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func createWorkerPool(noOfWorkers int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
            for job := range jobs {
                result := Result{job, digits(job.randomno), i} //create new struct
                results <- output
            }
        }(i)
	}
	wg.Wait()
	close(results)
}

func createParallelWorkers(noOfWorkers int, jobs chan Job) chan Result{
	var wg sync.WaitGroup
	results := make(chan Result, 10)
	go func(){
		defer close(results)
		defer wg.Wait()
		for i := 0; i < noOfWorkers; i++ {
			wg.Add(1)
						go func(i int) {
							defer wg.Done()
							for job := range jobs {
								output := Result{job, digits(job.randomno), i} //create new struct
								results <- output
							}
						}(i)
					}
				}()
				return results
	}


func allocate(noOfJobs int, jobs chan Job){
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999) + 1
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func allocateParallel(noOfJobs int) chan Job {
	jobs := make(chan Job, 10)
	go func(){
		defer close(jobs)
        for i := 0; i < noOfJobs; i++ {
            randomno := rand.Intn(999) + 1
            job := Job{i, randomno}
            jobs <- job
        }
	}()
	return jobs
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	noOfJobs := 50
	go allocate(noOfJobs, jobs)
	noOfWorkers := 10
	go createWorkerPool(noOfWorkers, jobs, results)

	for result := range results {
        fmt.Println(result.job.id, result.job.randomno, result.sumofdigits, result.wId)
    }
}