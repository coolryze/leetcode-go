package implementstackusingqueues

import "testing"

func TestMyStack(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)
	if stack.Empty() {
		t.Errorf("Push failed")
	}

	if stack.Top() != 2 {
		t.Errorf("Peek failed")
	}

	if stack.Pop() != 2 {
		t.Errorf("Pop failed")
	}

	if stack.Empty() {
		t.Errorf("Empty failed")
	}
}
