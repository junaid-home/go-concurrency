package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var count int32 = 12

func main() {
	fmt.Printf("Initial Count: %d\n", count)
	wg.Add(2)

	go doInc()
	go doDec()

	wg.Wait()
	fmt.Printf("Final Count: %d\n", count)
}

func doInc() {
	for i := 0; i < 30000; i++ {
		atomic.AddInt32(&count, 1)
	}
	wg.Done()
}

func doDec() {
	for i := 0; i < 30000; i++ {
		atomic.AddInt32(&count, -1)
	}
	wg.Done()
}
