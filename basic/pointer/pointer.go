package main

import "fmt"

func main() {
	modifyValueByPointer()
	var a int = 10
	passByValue(a)
	fmt.Printf("pass by value, result: %d\n", a)
	passByPointer(&a)
	fmt.Printf("pass by pointer, result: %d\n", a)
	fmt.Println("swap:")
	a, b := 3, 4
	fmt.Println(swap(a, b))
}

func modifyValueByPointer() {
	var a int = 2
	pa := &a
	*pa = 3
	fmt.Println(a)
}

func passByValue(a int) {
	a++
}

func passByPointer(a *int) {
	*a++
}

func swap(a, b int) (int, int) {
	return b, a
}
