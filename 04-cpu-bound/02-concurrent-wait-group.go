package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(4)
	start := time.Now()

	wg.Add(4)

	go runA()
	go runB()
	go runC()
	go runD()

	wg.Wait()

	total := time.Since(start)

	fmt.Printf("Total Execution Time: %s", total)
}

func runA() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	fmt.Printf("Total Count A: %d\n", count)
	wg.Done()
}

func runB() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	fmt.Printf("Total Count B: %d\n", count)
	wg.Done()
}

func runC() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	fmt.Printf("Total Count C: %d\n", count)
	wg.Done()
}

func runD() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	fmt.Printf("Total Count D: %d\n", count)
	wg.Done()
}
