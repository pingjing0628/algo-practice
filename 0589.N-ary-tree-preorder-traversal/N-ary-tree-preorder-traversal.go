package problem0589

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, result := []*Node{root}, []int{}

	for len(stack) > 0{
		pop := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		result = append(result, pop.Val)

		for i := len(pop.Children) - 1; i >= 0; i-- {
			stack = append(stack, pop.Children[i])
		}
	}
	return result
}

