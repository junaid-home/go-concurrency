package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	end := make(chan string)

	dostuff := func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Calculating...")

		ch <- "230 kg"
	}

	printResult := func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Result: " + <-ch)
		end <- "Exit"
	}

	go dostuff()
	go printResult()

	<-end
}
