package main

import "fmt"

func main() {
	partition(100)
	arr := []int{}
	for i := 1; i <= 100; i++ {
		arr = append(arr, len(PART[i]))
	}
	fmt.Println(arr)
}

var PART = map[int][][]int{
	0: {},
	1: {[]int{1}},
}

func partition(n int) [][]int {
	if k, ok := PART[n]; ok {
		return k
	}
	res := make([][]int, 0)
	res = append(res, []int{n})
	for i := n; i > 0; i-- {
		arr := partition(n - i)
		for _, v := range arr {
			if v[0] <= i {
				t := append([]int{i}, v...)
				res = append(res, t)
			}
		}
	}
	PART[n] = res
	fmt.Println(n, len(res))
	return res
}
