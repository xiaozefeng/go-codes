package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else{
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	//tryRecover()

	//b:=0
	//a:= 1/b
	//fmt.Println(a)

	panic(123)
}

func tryRecover() {
	panic(errors.New("this is an error"))
}
