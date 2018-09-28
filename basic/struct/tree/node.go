package tree

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func CreateNode(val int) *Node {
	return &Node{Val: val}
}

func (t *Node) Traverse() {
	if t == nil {
		return
	}
	t.Left.Traverse()
	t.Print()
	t.Right.Traverse()
}

func (t Node) Print() {
	fmt.Print(t.Val, " ")
}

func (t *Node) SetValue(val int) {
	t.Val = val
}

func (t *Node) String() string {
	return strconv.Itoa(t.Val)
}
