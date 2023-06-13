package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct { 	
	URL, workerIdMsg    string
	Status int
}

func pingWebsite(wid int, jobs <- chan Site, results chan <- Result) {
	for site := range jobs {
        log.Printf("Pinging %s\n", site.URL)
        resp, err := http.Get(site.URL)
        if err!= nil {
            log.Println(err.Error())
            
        }
        results <- Result{
            URL: site.URL,
            workerIdMsg: fmt.Sprintf("Pong %s", site.URL),
            Status: resp.StatusCode,
        }
    }
}

func main(){
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 4; w++ {
		go pingWebsite(w, jobs, results)
	}

	urls := []string{
		"https://educative.io",
		"https://educative.io/learn",
		"https://educative.io/teach",
		"https://www.educative.io/explore/new",
		"https://www.educative.io/explore/picks",
		"https://www.educative.io/explore/early-access",
		"https://google.com",
	}

	for _, url := range urls {
        jobs <- Site{url}
    }
	close(jobs)

	for i := 1; i <= len(urls); i++ {
		result := <- results
        log.Printf("Result: %s %d\n", result.URL, result.Status)
	}
}