package partitionlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		name string
		x    int
		head []int
		res  []int
	}{
		{"ok", 3, []int{1, 4, 3, 2, 5, 2}, []int{1, 2, 2, 4, 3, 5}},
		{"2 elems", 2, []int{2, 1}, []int{1, 2}},
	}

	for _, c := range cases {
		res := partition(createListFromArray(c.head), c.x)
		require.Equal(t, getArrayFromList(res), c.res)
	}
}

func createListFromArray(nums []int) *ListNode {
	l := &ListNode{}

	t := l
	for i := range nums {
		t.Next = &ListNode{Val: nums[i]}
		t = t.Next
	}

	return l.Next
}

func getArrayFromList(l *ListNode) []int {
	var arr []int
	for t := l; t != nil; t = t.Next {
		arr = append(arr, t.Val)
	}

	return arr
}
