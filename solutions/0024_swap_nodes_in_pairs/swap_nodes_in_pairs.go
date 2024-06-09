package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		temp1 := cur.Next
		temp2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = temp1
		temp1.Next = temp2
		cur = cur.Next.Next
	}

	return dummyHead.Next
}
