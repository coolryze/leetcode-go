package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy
	step := 0

	for fast.Next != nil {
		fast = fast.Next
		step++
		if step > n {
			slow = slow.Next
		}
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
