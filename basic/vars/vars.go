package main

import (
	"fmt"
	"math"
)

// package internal variable
var (
	pi      = 1.1314
	version = "1.0"
	flag    = false
)

// variable zero value
func variableZeroValue() {
	var s string
	var a int
	fmt.Printf("%q, %d\n", s, a)
}

// variable initial value
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("a = %d, b = %d ,s = %s\n", a, b, s)
}

// variable type deduction
func variableTypeDeduction() {
	var a, b, c = 3, 4, true
	var s = "abc"
	fmt.Printf("a = %d, b = %d , c =%t ,s = %s\n", a, b, c, s)
}

// variable shorter
func variableShorter() {
	a, b, c := 1, 2, false
	s := "abc"
	fmt.Printf("a = %d, b = %d , c =%t, s = %s\n", a, b, c, s)
}

// calculate triangle
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + a*b)))
	fmt.Println(c)
}

// contants
func contants() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*b + b*b))
	fmt.Printf("filename=%s, c=%d\n", filename, c)
}

// enums
func enums() {
	const (
		golang = iota
		java
		python
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Printf("golang=%d, java=%d, python=%d\n", golang, java, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("variableZeroValue:")
	variableZeroValue()
	fmt.Println("variableInitialValue:")
	variableInitialValue()
	fmt.Println("variableTypeDeduction:")
	variableTypeDeduction()
	fmt.Println("variableShorter:")
	variableShorter()
	fmt.Println("package internal variables: ")
	fmt.Printf("%f, %t, %s\n", pi, flag, version)
	fmt.Println("triangle:")
	triangle()
	fmt.Println("constants:")
	contants()
	fmt.Println("enums")
	enums()
}
