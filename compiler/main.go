package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	l := 1
	r := n
	mid := (l + r) / 2
	for l <= r {
		mid = (l + r) / 2
		fmt.Println(mid)

		var ans int
		fmt.Scan(&ans)
		if ans == 0 {
			r = mid
		} else {
			l = mid + 1
		}
		if l == r {
			fmt.Println("!", l)
			return
		}
	}
}
