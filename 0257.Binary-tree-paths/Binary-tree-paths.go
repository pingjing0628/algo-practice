package problem0257

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	output := []string{}

	var path = func(*TreeNode, string) {}

	path = func(root *TreeNode, s string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			fmt.Println(root.Val)
			fmt.Println(root.Left)
			output = append(output, s+strconv.Itoa(root.Val))
			return
		}

		s += strconv.Itoa(root.Val) + "->"

		path(root.Left, s)
		path(root.Right, s)
	}

	path(root, "")
	return output
}
