package designlinkedlist

import "testing"

func TestMyLinkedList(t *testing.T) {
	linkedList := Constructor()

	linkedList.AddAtHead(1)
	if linkedList.Get(0) != 1 {
		t.Errorf("AddAtHead failed")
	}

	linkedList.AddAtTail(3)
	if linkedList.Get(1) != 3 {
		t.Errorf("AddAtTail failed")
	}

	linkedList.AddAtIndex(1, 2)
	if linkedList.Get(1) != 2 {
		t.Errorf("AddAtIndex failed")
	}

	linkedList.DeleteAtIndex(1)
	if linkedList.Get(1) != 3 {
		t.Errorf("DeleteAtIndex failed")
	}
}
