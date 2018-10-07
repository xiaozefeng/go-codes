package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel demo:")
	channelDemo()

	fmt.Println("buffered channel:")
	bufferedChannel()

	fmt.Println("range and close channel:")
	rangeAndCloseChannel()
	fmt.Println()

	fmt.Println("select demo:")
	selectDemo()

	fmt.Println()
	fmt.Println("default select demo:")
	defaultSelect()

	fmt.Println()
	fmt.Println("exercise binary tree:")
	binaryTreeExercise()
}

func defaultSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func selectDemo() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d, ", <-c)
		}
		quit <- 0
	}()
	fib(c, quit)
}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println()
			fmt.Println("quit")
			return
		}
	}
}

func rangeAndCloseChannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("%d, ", i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func bufferedChannel() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 2
	ch <- 1
	fmt.Printf("%d, %d, %d\n", <-ch, <-ch, <-ch)
}

func channelDemo() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("%d, %d, %d", x, y, x+y)
}

// sum
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func binaryTreeExercise() {
	root := &Tree{Val: 3}
	root.Left = &Tree{Val: 1}
	root.Left.Right = &Tree{Val: 2}
	root.Right = &Tree{Val: 5}

	c := make(chan int)
	go walk(root, c)
	for i := range c {
		fmt.Printf("%d ", i)
	}


	result := same(&Tree{Val: 3}, &Tree{Val: 3})
	fmt.Printf("eqauls result: %t", result)
}

type Tree struct {
	Val         int
	Left, Right *Tree
}

func newTree(val int) *Tree {
	return &Tree{
		Val: val,
	}
}

func walk(t *Tree, c chan int) {
	t.traverse(func(t *Tree) {
		c <- t.Val
	})
	close(c)
}

func (t *Tree) traverse(f func(t *Tree)) {
	if t == nil {
		return
	}
	t.Left.traverse(f)
	f(t)
	t.Right.traverse(f)
}

func same(t1 *Tree, t2 *Tree) bool {
	return t1.Val == t2.Val
}
