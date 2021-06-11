package problem1290

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	current := head
	result := 0
	for current != nil {
		result = result << 1
		result += current.Val
		current = current.Next
	}

	return result
}
