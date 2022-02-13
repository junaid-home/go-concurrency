package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 1)
			ch1 <- "[ CH 1 ] Sending After 1 Seconds"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "[ CH 2 ] Sending After 2 Seconds"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			ch3 <- "[ CH 3 ] Sending After 10 Seconds"
		}
	}()

	for {
		select {
		case main := <-ch1:
			fmt.Printf("Received from channel 1: %s\n", main)
		case main := <-ch2:
			fmt.Printf("Received from channel 2: %s\n", main)
		case main := <-ch3:
			fmt.Printf("Received from channel 3: %s\n", main)
			os.Exit(1)
		}
	}
}
