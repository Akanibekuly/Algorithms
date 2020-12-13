package main

import (
	"fmt"
)

func main() {
	var n, x, k int
	reminders := map[int]int{}
	fmt.Scan(&n, &x, &k)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		rem := temp % x
		v, k := reminders[rem]
		if !k {
			reminders[rem] = temp
		} else if temp < v {
			reminders[rem] = temp
		}
	}

	getTotal := func(t int) int {
		sum := 0
		for _, v := range reminders {
			temp := (t + 1 - v) / x
			if temp >= 0 {
				sum += temp + 1
			}
		}
		return sum
	}

	l := 0
	r := 1000000000
	for l < r {
		m := (l + r) / 2
		total := getTotal(m)
		if total < k {
			l = m + 1
		} else {
			r = m
		}
		if l == r {
			fmt.Println(l)
			return
		}
	}
}
