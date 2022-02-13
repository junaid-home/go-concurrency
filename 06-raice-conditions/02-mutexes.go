package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var count int = 12

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
		mutex.Lock()
		count -= 1
		mutex.Unlock()
	}
	wg.Done()
}

func doDec() {
	for i := 0; i < 30000; i++ {
		mutex.Lock()
		count += 1
		mutex.Unlock()
	}
	wg.Done()
}
