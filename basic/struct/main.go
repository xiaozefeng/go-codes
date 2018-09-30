package main

import (
	"fmt"

	"github.com/xiaozefeng/go-codes/basic/struct/queue"
	"github.com/xiaozefeng/go-codes/basic/struct/tree"
)

func main() {
	fmt.Println("About Tree Node")
	aboutTreeNode()
	fmt.Println("About Queue")
	aboutQueue()
}
func aboutQueue() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
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
	fmt.Println("PostOrder()")
	m := myTreeNode{&root}
	m.postOrder()
	fmt.Println()

	fmt.Println("root print")
	root.Print()
	root.SetValue(100)
	root.Print()
	fmt.Println()
}

type myTreeNode struct {
	node *tree.Node
}

func (my *myTreeNode) postOrder() {
	if my == nil || my.node == nil {
		return
	}
	left := myTreeNode{my.node.Left}
	left.postOrder()
	right := myTreeNode{my.node.Right}
	right.postOrder()
	my.node.Print()
}
