package partitionlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	fhead, shead := &ListNode{}, &ListNode{}
	ftail, stail := fhead, shead
	for t := head; t != nil; t = t.Next {
		if t.Val < x {
			ftail.Next = &ListNode{Val: t.Val}
			ftail = ftail.Next
		} else {
			stail.Next = &ListNode{Val: t.Val}
			stail = stail.Next
		}
	}

	ftail.Next = shead.Next

	return fhead.Next
}

func (l *ListNode) print() {
	if l == nil {
		fmt.Println("(nil)")
		return
	}

	for t := l; t != nil; t = t.Next {
		fmt.Print(t.Val)
		if t.Next != nil {
			fmt.Print("->")
		} else {
			fmt.Println()
		}
	}
}
