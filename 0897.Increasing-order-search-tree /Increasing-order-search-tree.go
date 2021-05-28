package problem0897

type TreeNode struct {
	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	min, _ := find(root)

	return min
}

func find(node *TreeNode) (*TreeNode, *TreeNode) {
	min, max := node, node

	if node.Left != nil {
		leftMin, leftMax := find(node.Left)
		min = leftMin
		leftMax.Right = node // 將 node 往右邊排
	}

	if node.Right != nil {
		rightMin, rightMax := find(node.Right)
		max = rightMax
		node.Right = rightMin // 將 node 原右側往後排
	}

	node.Left = nil // 左邊值需補上 null
	return min, max
}