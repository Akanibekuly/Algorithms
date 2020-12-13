package main

import "fmt"

// storage of factorials
var FACT = map[int]int{0: 1}

// storage of numbers partitions
var PART = map[int][][]int{
	0: {},
	1: {[]int{1}},
}

func main() {
	fmt.Println(Relationship(4))
}

func Relationship(n int) int {
	if n == 0 {
		return 1
	}
	factorial(n)
	p := partition(n)
	res := 0
	for _, v := range p {
		prod := 1
		t := n
		// fmt.Println(v)
		uniq := make(map[int]int)
		for _, k := range v {
			uniq[k]++
			cnk := FACT[t] / (FACT[k] * FACT[t-k])
			prod *= cnk
			t -= k
		}
		prod *= FACT[len(v)]
		// fmt.Println(i, v, uniq, prod)
		for _, v := range uniq {
			prod /= FACT[v]
		}
		// fmt.Println(prod)
		res += prod
	}
	return res
}

func factorial(n int) int {
	if k, ok := FACT[n]; ok {
		return k
	}
	res := n * factorial(n-1)
	FACT[n] = res
	return res
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
	return res
}
