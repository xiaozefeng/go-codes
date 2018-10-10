package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel demo:")
	channelDemo()
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("Workder: %d received %c\n", id, <-c)
	}
}

func channelDemo() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go worker(i, c)
	}

	for i := 0; i < 10; i++ {
		c <- 'a' + i
	}

	fmt.Println()
	for i := 0; i < 10; i++ {
		c <- 'A' + i
	}
	time.Sleep(1 * time.Millisecond)
}
