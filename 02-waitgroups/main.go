package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()

	wg.Add(2) // we have 2 go routines

	go firstBlockingTask()
	go secondBlockingTask()

	wg.Wait() // wait for go routines to finish

	elapsed := time.Since(start)

	fmt.Printf("Total Time: %s", elapsed)
}

func firstBlockingTask() {
	time.Sleep(time.Second * 2)
	fmt.Println("Finished First Blocking Task")
	wg.Done()
}

func secondBlockingTask() {
	time.Sleep(time.Second * 2)
	fmt.Println("Finished Second Blocking Task")
	wg.Done()
}
