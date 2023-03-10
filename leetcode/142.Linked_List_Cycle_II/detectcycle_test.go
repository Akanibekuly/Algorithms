package linkedlistcycleii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	ex1Tail := &ListNode{Val: -4}
	ex1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val:  0,
				Next: ex1Tail,
			}}}
	ex1Tail.Next = ex1.Next

	ex2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	ex2.Next.Next = ex2

	cases := []struct {
		head *ListNode
		res  *ListNode
	}{
		{&ListNode{Val: 1}, nil},
		{ex1, ex1.Next},
		{ex2, ex2},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, detectCycle(c.head))
		})
	}
}
