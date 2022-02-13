package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	go func() {
		fmt.Println("Putting to the buffer...")
		ch <- "First Value "
		fmt.Println("Putting to the buffer...")
		ch <- "Second Value "
		fmt.Println("Putting to the buffer...")
		ch <- "Third Value"
		fmt.Println("buffer is full...")
	}()

	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
}
