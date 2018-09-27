package main

import (
	"fmt"
	"runtime"
	"time"
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

func printOS() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("osx")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

func main() {
	printOS()
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	log.Printf("%s", err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	fmt.Println("eval:")
	r1 := eval(1, 2, "+")
	r2 := eval(1, 2, "-")
	r3 := eval(1, 2, "*")
	r4 := eval(1, 2, "/")
	fmt.Println(r1, r2, r3, r4)
	printToday()
	printTime()
}

func printTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func printToday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In tow days")
	default:
		fmt.Println("Too far away")
	}
}
