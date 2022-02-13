package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	start := time.Now()

	links := []string{"https://www.bing.com/", "https://www.yahoo.com/", "https://duckduckgo.com/", "https://www.google.com/"}

	wg.Add(len(links))
	for _, l := range links {
		go checkConn(l)
	}

	wg.Wait()

	total := time.Since(start)

	fmt.Printf("Total Execution Time: %s", total)
}

func checkConn(url string) {
	if _, err := http.Get(url); err != nil {
		fmt.Printf("[ DOWN ]: %s\n", url)
		return
	}

	fmt.Printf("[ UP ]: %s\n", url)
	wg.Done()
}
