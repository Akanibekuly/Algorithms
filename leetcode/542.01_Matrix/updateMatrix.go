package matrix

func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	if rows == 0 {
		return mat
	}

	cols := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				mat[i][j] = 10000
				if i > 0 {
					min(&mat[i][j], mat[i-1][j]+1)
				}
				if j > 0 {
					min(&mat[i][j], mat[i][j-1]+1)
				}
			}
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if i < rows-1 {
					min(&mat[i][j], mat[i+1][j]+1)
				}
				if j < cols-1 {
					min(&mat[i][j], mat[i][j+1]+1)
				}
			}
		}
	}

	return mat
}

func min(a *int, b int) {
	if *a < b {
		*a = b
	}
}

// func updateMatrix(mat [][]int) [][]int {

// 	rows := len(mat)
// 	if rows == 0 {
// 		return mat
// 	}
// 	cols := len(mat[0])

// 	//First check
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if mat[i][j] == 1 {
// 				mat[i][j] = 10000
// 				if i > 0 {
// 					min(&mat[i][j], mat[i-1][j]+1)
// 				}
// 				if j > 0 {
// 					min(&mat[i][j], mat[i][j-1]+1)
// 				}
// 			}
// 		}
// 	}

// 	//Second check
// 	for i := rows - 1; i >= 0; i-- {
// 		for j := cols - 1; j >= 0; j-- {
// 			if i < rows-1 {
// 				min(&mat[i][j], mat[i+1][j]+1)
// 			}
// 			if j < cols-1 {
// 				min(&mat[i][j], mat[i][j+1]+1)
// 			}
// 		}
// 	}

// 	return mat
// }

// func min(d1 *int, d2 int) {
// 	if *d1 > d2 {
// 		*d1 = d2
// 	}
// }
