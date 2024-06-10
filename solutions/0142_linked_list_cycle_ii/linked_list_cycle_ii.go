package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			index1 := fast
			index2 := head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}

	return nil
}
