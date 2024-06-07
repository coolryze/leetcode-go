package removelinkedlistelements

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
	val  int
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

func TestRemoveElements(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 6, 3, 4, 5, 6}, 6},
			output{[]int{1, 2, 3, 4, 5}},
		},
		testcase{
			input{[]int{}, 1},
			output{[]int{}},
		},
		testcase{
			input{[]int{7, 7, 7, 7}, 7},
			output{[]int{}},
		},
	}

	for _, tc := range tcs {
		got := linkedListToArray(removeElements(arrayToLinkedList(tc.input.head), tc.input.val))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
