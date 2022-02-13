package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	go firstBlockingTask()
	go secondBlockingTask()

	time.Sleep(time.Second * 3)

	elapsed := time.Since(start)

	fmt.Printf("Total Time: %s", elapsed)
}

func firstBlockingTask() {
	time.Sleep(time.Second * 2)
	fmt.Println("Finished First Blocking Task")
}

func secondBlockingTask() {
	time.Sleep(time.Second * 2)
	fmt.Println("Finished Second Blocking Task")
}
