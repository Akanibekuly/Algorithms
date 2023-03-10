package linkedlistrandomnode

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	arr []*ListNode
}

func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().Unix())

	var s Solution
	for t := head; t != nil; t = t.Next {
		s.arr = append(s.arr, t)
	}

	return s
}

func (this *Solution) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))].Val
}
