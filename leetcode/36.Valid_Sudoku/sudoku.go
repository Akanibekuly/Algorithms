package validsudoku

func isValidSudoku(board [][]byte) bool {

	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				if !isValid(board, i, j) {
					return false
				}
			}
		}
	}
	return true
}

func isValid(board [][]byte, i, j int) bool {
	for m := 0; m < len(board); m++ {
		if m != i && board[i][j] == board[m][j] {
			return false
		}
		if m != j && board[i][j] == board[i][m] {
			return false
		}
	}
	x, y := i-i%3, j-j%3
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			if x+m != i && y+n != j && board[x+m][y+n] == board[i][j] {
				return false
			}
		}
	}
	return true
}
