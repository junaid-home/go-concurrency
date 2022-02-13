package main

import (
	"fmt"
	"runtime"
	"time"
)

var ch = make(chan int)
var end = make(chan string)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()

	go runA()
	go runB()
	go runC()
	go runD()

	for i := 0; i < 4; i++ {
		fmt.Printf("Total Count %d: %d\n", i, <-ch)
	}

	total := time.Since(start)

	fmt.Printf("Total Execution Time: %s", total)
	<-end
}

func runA() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	ch <- count
}

func runB() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	ch <- count
}

func runC() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	ch <- count
}

func runD() {
	var count = 0
	for i := 1; i < 10000000000; i++ {
		count += i
	}

	ch <- count
	end <- "Exit"
}
