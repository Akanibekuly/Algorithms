package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for i := range nums {
		dict[nums[i]]++
	}
	arr := make([]int, 0, len(dict))
	for k := range dict {
		arr = append(arr, k)
	}
	//fmt.Println(dict, arr)
	sort.Slice(arr, func(i, j int) bool {
		return dict[arr[i]] > dict[arr[j]]
	})
	//fmt.Println(arr)

	return arr[:k]
}

func topKFrequentHeap(nums []int, k int) []int {
	m := maxHeap{
		nums: make([]int, 0, len(nums)),
		dict: make(map[int]int, len(nums)),
	}

	heap.Init(&m)

	for i := range nums {
		heap.Push(&m, nums[i])
	}

	fmt.Println(m)

	return m.nums[:k]
}

type maxHeap struct {
	nums []int
	dict map[int]int
}

func (m *maxHeap) Less(i, j int) bool {
	return m.dict[m.nums[i]] > m.dict[m.nums[j]]
}

func (m *maxHeap) Len() int {
	return len(m.nums)
}

func (m *maxHeap) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}

func (m *maxHeap) Push(v interface{}) {
	d := v.(int)
	if _, ok := m.dict[d]; !ok {
		m.nums = append(m.nums, d)
	}
	m.dict[d]++
}

func (m *maxHeap) Pop() (v interface{}) {
	v, m.nums = m.nums[m.Len()-1], m.nums[:m.Len()-1]
	if _, ok := m.dict[v.(int)]; ok {
		m.dict[v.(int)]--
		if m.dict[v.(int)] == 0 {
			delete(m.dict, v.(int))
		}
	}
	return
}
