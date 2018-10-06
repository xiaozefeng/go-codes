package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.inc("somekey")
	}
	time.Sleep(1 *time.Second)
	fmt.Println(c.value("somekey"))

}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *SafeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
