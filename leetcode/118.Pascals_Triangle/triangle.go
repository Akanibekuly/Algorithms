package pascal

var rows = map[int][]int{0: {1}, 1: {1, 1}}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := range res {
		row, ok := rows[i]
		if ok {
			res[i] = row
		} else {
			res[i] = make([]int, i+1)
			res[i][0], res[i][i] = 1, 1
			for j := 1; j < i; j++ {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}

			rows[i] = res[i]
		}
	}

	return res
}
