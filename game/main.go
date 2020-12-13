package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	fmt.Println(k, n)
	countA := 0
	countB := 0
	for i := 0; i <= n; i++ {
		var temp int
		fmt.Scan(&temp)
		fmt.Println(temp)
		if temp%3 == 0 && temp%5 == 0 {
			continue
		} else if temp%3 == 0 {
			countA++
			if countA == k {
				fmt.Println("Petya")
				return
			}
		} else if temp%5 == 0 {
			countB++
			if countB == k {
				fmt.Println("Vasya")
				return
			}
		}
	}
	fmt.Println("Draw")
}
