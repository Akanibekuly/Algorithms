package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hare, tortoise := head, head
	cycle := false
	for hare.Next != nil && hare.Next.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next

		if hare == tortoise {
			cycle = true
			break
		}
	}

	if !cycle {
		return nil
	}

	tortoise = head
	for tortoise != hare {
		hare, tortoise = hare.Next, tortoise.Next
	}

	return tortoise
}
