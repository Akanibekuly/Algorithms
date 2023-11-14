package seatreservationmanager

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Push(v any) { *h = append(*h, v.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type SeatManager struct {
	hq     IntHeap
	idx, m int
}

func Constructor(n int) SeatManager {
	hq := IntHeap{}
	heap.Init(&hq)

	return SeatManager{
		hq: hq,
		m:  n,
	}
}

func (s *SeatManager) Reserve() int {
	if s.hq.Len() > 0 {
		return heap.Pop(&s.hq).(int)
	}
	s.idx++
	return s.idx
}

func (s *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&s.hq, seatNumber)
}
