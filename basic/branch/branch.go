package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const filename = "basic/vars/README.MD"

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func main() {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		log.Printf("%s", err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println("eval:")
	r1 := eval(1, 2, "+")
	r2 := eval(1, 2, "-")
	r3 := eval(1, 2, "*")
	r4 := eval(1, 2, "/")
	fmt.Println(r1, r2, r3, r4)
}
