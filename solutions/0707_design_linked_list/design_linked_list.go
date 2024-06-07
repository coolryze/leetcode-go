package designlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	dummyHead *ListNode
	Size      int
}

func Constructor() MyLinkedList {
	newNode := &ListNode{
		-1,
		nil,
	}
	return MyLinkedList{
		dummyHead: newNode,
		Size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val}
	node.Next = this.dummyHead.Next
	this.dummyHead.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{Val: val}
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	node := &ListNode{Val: val}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}
