package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com",
	"https://www.somewhere.com",
	"https://www.yahoo.com",
	"https://www.someplace.com",
}

func fetchUrl(url string, wg *sync.WaitGroup){
	defer wg.Done() //decrement counter
    resp, err := http.Get(url)
    if err!= nil {
        fmt.Println(err)
       
    }
    
    fmt.Println(resp.Status)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	var wg sync.WaitGroup
	fmt.Println("Home endpoint")
	wg.Add(len(urls))
	for _, url := range urls {
        go fetchUrl(url, &wg)
    }
	wg.Wait()
	fmt.Fprintf(w, "all responses received successfully")

}

func handleRequests() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
    handleRequests()
}