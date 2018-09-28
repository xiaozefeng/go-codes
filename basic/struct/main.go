package main

import (
	"fmt"

	"github.com/xiaozefeng/go-codes/basic/struct/tree"
)

func main() {
	aboutTreeNode()
}

func aboutTreeNode() {
	root := tree.Node{Val: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	fmt.Println()
	fmt.Println("Node Traverse():")
	root.Traverse()
	fmt.Println()

	fmt.Println("root print")
	root.Print()
	root.SetValue(100)
	root.Print()
	fmt.Println()
}
