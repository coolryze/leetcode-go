package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := head
	for cur != nil {
		newHead.Next = &ListNode{cur.Val, newHead.Next}
		cur = cur.Next
	}
	return newHead.Next
}

func reverseList3(head *ListNode) *ListNode {
	return reverseList3Recursive(nil, head)
}

func reverseList3Recursive(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return reverseList3Recursive(head, next)
}
