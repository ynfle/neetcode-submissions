func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		count := buildRowCount(board, row)
		if !countIsValid(count) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		count := buildColCount(board, col)
		if !countIsValid(count) {
			return false
		}
	}

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			count := buildBoxCount(board, row, col)
			if !countIsValid(count) {
				return false
			}
		}
	}
	return true
}

func countIsValid(count [9]int) bool {
	for _, c := range count {
		if c > 1 {
			return false
		}
	}
	return true
}

func buildRowCount(board [][]byte, row int) [9]int {
	count := [9]int{}
	for _, b := range board[row] {
		if b != '.' {
			count[b-'1'] += 1
		}
	}
	return count
}

func buildColCount(board [][]byte, col int) [9]int {
	count := [9]int{}
	for row := 0; row < 9; row++ {
		b := board[row][col]
		if b != '.' {
			count[b-'1'] += 1
		}
	}
	return count
}

func buildBoxCount(board [][]byte, boxRow int, boxCol int) [9]int {
	count := [9]int{}
	for row := boxRow*3; row < boxRow*3+3; row++ {
		for col := boxCol*3; col < boxCol*3+3; col++ {
			b := board[row][col]
			if b != '.' {
				count[b-'1'] += 1
			}
		}
	}
	return count
}