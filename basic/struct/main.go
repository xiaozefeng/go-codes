package main

import (
	"fmt"

	"github.com/xiaozefeng/go-codes/basic/struct/tree"
)

func main() {
	aboutTreeNode()
}

func aboutTreeNode() {
	root := tree.TreeNode{Val: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	fmt.Println(root)

	nodes := []tree.TreeNode{
		{3, nil, nil},
		{5, nil, nil},
		{7, nil, &root},
	}
	fmt.Println(nodes)
}
