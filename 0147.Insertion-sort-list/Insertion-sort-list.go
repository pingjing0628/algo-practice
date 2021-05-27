package problem0147

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	cur := head.Next

	head.Next = nil
	for cur != nil {

		node := cur
		cur = cur.Next
		node.Next = nil
		if node.Val > tail.Val {
			tail.Next = node
			tail = node
			continue
		}
		if node.Val < head.Val {
			node.Next = head
			head = node
			continue
		}
		prev := head
		for prev.Next != nil && prev.Next.Val < node.Val {
			prev = prev.Next
		}
		node.Next = prev.Next
		prev.Next = node
	}
	return head

}
