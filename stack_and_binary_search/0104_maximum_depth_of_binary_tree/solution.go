package maximumdepthofbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}
