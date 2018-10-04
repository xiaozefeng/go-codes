package fib

import (
	"fmt"
	"io"
	"strings"
)

func Fibonacci() FibonacciGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//FibonacciGen
type FibonacciGen func() int

// 为函数实现接口
func (f FibonacciGen) Read(p []byte) (n int, err error) {
	next := f()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// 使用string帮我们做Read这个动作
	return strings.NewReader(s).Read(p)
}
