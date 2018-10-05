package main

import "fmt"

func main() {
	channelDemo()
}

func channelDemo() {
	//  c == nil
	// var c chan int
	c := make(chan int)
	go func() {
		fmt.Println(<-c)
	}()
	c <- 2

}
