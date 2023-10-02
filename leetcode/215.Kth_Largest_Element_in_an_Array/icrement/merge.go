package main

import "fmt"

func main() {
	s, t := []int{1, 3, 5}, []int{2, 4}

	fmt.Println(merge(s, t))
}

func merge(s, t []int) []int {
	res := make([]int, 0, len(s)+len(t))
	idx1, idx2 := 0, 0
	for idx1 < len(s) && idx2 < len(t) {
		if s[idx1] > t[idx2] {
			res = append(res, t[idx2])
			idx2++
		} else {
			res = append(res, s[idx1])
			idx1++
		}
	}
	if idx1 < len(s) {
		res = append(res, s[idx1:]...)
	}
	if idx2 < len(t) {
		res = append(res, t[idx2:]...)
	}

	return res
}
