package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱编程"
	fmt.Printf("我爱编程length:%d\n", len(s))

	for i, b := range []byte(s) {
		fmt.Printf("(%d, %X)", i, b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d,%X)", i, ch)
	}
	fmt.Println()
	fmt.Printf("Rune Count int string: %d\n", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", r)
	}
	fmt.Println()

	for i, val := range []rune(s) {
		fmt.Printf("(%d, %c)", i, val)
	}
	fmt.Println()

}
