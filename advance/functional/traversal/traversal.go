package traversal

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (n *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}
