package linkedlistrandomnode

import (
	"math/rand"
	"time"
)

type Solution2 struct {
	arr []*ListNode
}

func Constructor2(head *ListNode) Solution {
	rand.Seed(time.Now().Unix())

	var s Solution
	for t := head; t != nil; t = t.Next {
		s.arr = append(s.arr, t)
	}

	return s
}

func (this *Solution2) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))].Val
}
