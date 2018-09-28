package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func CreateNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
