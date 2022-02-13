package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()

	go runA()
	go runB()
	go runC()
	go runD()

	total := time.Since(start)

	fmt.Printf("Total Execution Time: %s", total)
}

func runA() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	fmt.Printf("Total Count: %d\n", count)
}

func runB() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}
	fmt.Printf("Total Count: %d\n", count)
}

func runC() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}
	fmt.Printf("Total Count: %d\n", count)
}

func runD() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}
	fmt.Printf("Total Count: %d\n", count)
}
