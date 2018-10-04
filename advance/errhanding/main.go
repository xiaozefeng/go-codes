package main

import (
	"bufio"
	"fmt"
	"github.com/xiaozefeng/go-codes/advance/errhanding/deferhanding"
	"github.com/xiaozefeng/go-codes/advance/functional/fib"
	"os"
)

func main() {
	fmt.Println("defer:")
	deferhanding.TryDefer()
	fmt.Println("write file:")
	writeFile("fibonacci.txt")
}

// write file
func writeFile(filename string) {
	file ,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}
