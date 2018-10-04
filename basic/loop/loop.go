package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(1))
	printFile("basic/loop/loop.go")
	fmt.Println()
	fmt.Println("print str contens:")
	str := `abc
	bcd
	fgh`
	printContents(strings.NewReader(str))

	fmt.Println()
	fmt.Println("for:")
	arr := make([]int, 10)
	// for i
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	//for range
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

}

// convert number to binary
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// print file content
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printContents(file)
}

func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
