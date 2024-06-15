package implementqueueusingstacks

import "testing"

func TestMyQueue(t *testing.T) {
	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	if queue.Empty() {
		t.Errorf("Push failed")
	}

	if queue.Peek() != 1 {
		t.Errorf("Peek failed")
	}

	if queue.Pop() != 1 {
		t.Errorf("Pop failed")
	}

	if queue.Empty() {
		t.Errorf("Empty failed")
	}
}
