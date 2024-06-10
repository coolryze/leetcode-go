package intersectionoftwolinkedlists

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	headA *ListNode
	headB *ListNode
}

type output struct {
	expected *ListNode
}

func TestGetIntersectionNode(t *testing.T) {
	intersectionNode := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	headA := &ListNode{4, &ListNode{1, intersectionNode}}
	headB := &ListNode{5, &ListNode{0, &ListNode{1, intersectionNode}}}

	tcs := []testcase{
		testcase{
			input{headA, headB},
			output{intersectionNode},
		},
	}

	for _, tc := range tcs {
		got := getIntersectionNode(tc.input.headA, tc.input.headB)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.expected, got)
		}
	}
}
