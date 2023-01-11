package sudoku

import "fmt"

type cell struct {
	Value       byte
	GuessValues map[byte]struct{}
}

func solveSudoku(board [][]byte) {
	s := make([][]cell, 9)
	for i := range s {
		s[i] = make([]cell, 9)
	}

	fillSolvingBoard(s, board)

	fmt.Println(s)
}

func fillSolvingBoard(s [][]cell, board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				s[i][j].GuessValues = map[byte]struct{}{
					1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {},
				}
			} else {
				s[i][j].Value = board[i][j]
			}

		}
	}
}
