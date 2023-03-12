package reshapethematrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	getCell := func(i, j int) int {
		size := i*c + j

		k, l := size/n, size%n
		if l == -1 {
			l = size - 1
		}

		return mat[k][l]
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = getCell(i, j)
		}
	}

	return res
}
