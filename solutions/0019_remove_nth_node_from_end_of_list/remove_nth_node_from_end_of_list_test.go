package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	head []int
	n    int
}

type output struct {
	expected []int
}

func linkedListToArray(head *ListNode) []int {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3, 4, 5}, 2},
			output{[]int{1, 2, 3, 5}},
		},
		testcase{
			input{[]int{1}, 1},
			output{[]int{}},
		},
		testcase{
			input{[]int{1, 2}, 1},
			output{[]int{1}},
		},
	}

	for _, tc := range tcs {
		got := linkedListToArray(removeNthFromEnd(arrayToLinkedList(tc.input.head), tc.input.n))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
