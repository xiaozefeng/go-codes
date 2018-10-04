package main

import (
	"bufio"
	"fmt"
	"github.com/xiaozefeng/go-codes/advance/functional/adder"
	"github.com/xiaozefeng/go-codes/advance/functional/fib"
	"github.com/xiaozefeng/go-codes/advance/functional/traversal"
	"io"
)

func main() {
	fmt.Println("adder:")
	a := adder.Adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", a(i))
	}
	fmt.Println()

	fmt.Println("fibonacci:")
	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", f())
	}

	fmt.Println()
	fmt.Println("fibonacci第二种方式:")
	printReader(fib.Fibonacci())
	fmt.Println()

	fmt.Println("使用函数式编程中序遍历二叉树:")
	count := 0
	root := &traversal.TreeNode{Val: 3}
	root.Left = &traversal.TreeNode{Val: 1}
	root.Right = &traversal.TreeNode{Val: 5}
	root.TraverseFunc(func(node *traversal.TreeNode) {
		count++
		fmt.Printf("%d, ", node.Val)
	})
	fmt.Println("二叉树的节点数量:", count)
}

func printReader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("%s, ", scanner.Text())
	}
}
