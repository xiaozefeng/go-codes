package deferhanding

import "fmt"

func TryDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
