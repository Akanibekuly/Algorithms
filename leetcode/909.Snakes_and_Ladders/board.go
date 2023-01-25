package snakesandladders

import "fmt"

func snakesAndLadders(board [][]int) int {
	n := len(board)
	dest := make([]int, n*n+1)
	for i := range dest {
		dest[i] = -1
	}

	dest[1] = 0
	queue := []int{1}
	count := 0
	for len(queue) != 0 {

		fmt.Println(queue, count)
		fmt.Println(dest[1:])
		next := []int{}

		for _, curr := range queue {
			for i := curr + 1; i <= min(curr+6, n*n); i++ {
				if dest[i] == -1 {
					dest[i] = dest[curr] + 1
					x, y := getCoordinates(n, i)
					val := board[x][y]
					if val == -1 {
						next = append(next, i)
					} else {
						if dest[val] == -1 {
							dest[val] = dest[i]
							next = append(next, val)
						}
					}
				}
			}
		}

		count++
		queue = next
	}

	return dest[n*n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getCoordinates(n, curr int) (x, y int) {
	x = n - 1 - (curr-1)/n

	if (curr-1)/n%2 == 0 { // -> 0, 1, 2
		y = curr%n - 1
		if y < 0 {
			y += n
		}
	} else {
		y = n - curr%n // ->  5, 4, 3
		if y >= n {
			y -= n
		}
	}

	return
}
