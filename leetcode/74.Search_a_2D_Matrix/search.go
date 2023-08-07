package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	lr, lc := len(matrix), len(matrix[0])
	l, r := 0, lr*lc

	var mid int
	for l < r {
		mid = (l + r) / 2
		if matrix[mid/lc][mid%lc] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	a, b := l/lc, l%lc
	if a <= lr-1 && b <= lc-1 {
		return matrix[a][b] == target
	}

	return false
}
