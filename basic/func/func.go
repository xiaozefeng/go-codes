package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println("eval:")
	if result, err := eval(3, 4, "-"); err != nil {
		log.Printf("%s", err)
		return
	} else {
		fmt.Println(result)
	}
	fmt.Println("div:")
	fmt.Println(div(13, 3))
	fmt.Println("apply:")
	// rewrite function math.Pow
	fmt.Println(apply(pow, 3, 4))
	// anonymous function
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println("variable parameter:")
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return -1, fmt.Errorf("unsupported operation:" + op)
	}
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// functional programming
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// variable parameter
func sum(numbers ...int) int {
	result := 0
	for _, val := range numbers {
		result += val
	}
	return result
}
