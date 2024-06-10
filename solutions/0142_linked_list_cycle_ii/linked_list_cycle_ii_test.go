package linkedlistcycleii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	head *ListNode
}

type output struct {
	expected *ListNode
}

func TestDetectCycle(t *testing.T) {
	lastNode := &ListNode{-4, nil}
	head := &ListNode{3, &ListNode{2, &ListNode{0, lastNode}}}
	cycleFirstNode := head.Next
	lastNode.Next = cycleFirstNode

	tcs := []testcase{
		testcase{
			input{head},
			output{cycleFirstNode},
		},
	}

	for _, tc := range tcs {
		got := detectCycle(tc.input.head)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.expected, got)
		}
	}
}
