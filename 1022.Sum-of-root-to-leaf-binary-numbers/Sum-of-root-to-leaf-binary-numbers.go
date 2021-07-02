package problem1022

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, current int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return current + root.Val
	}

	current = (root.Val + current) << 1

	return dfs(root.Left, current) + dfs(root.Right, current)
}
