package mergeksortedlists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {

	cases := []struct {
		lists []*ListNode
		res   *ListNode
	}{
		{
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 3}},
				{Val: 2, Next: &ListNode{Val: 4}},
			},
			res: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						}}}},
		},
		{
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				{Val: 2, Next: &ListNode{Val: 6}},
			},
			res: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 4, Next: &ListNode{
										Val: 5, Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{lists: nil, res: nil},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, mergeKLists(c.lists))
		})
	}
}
