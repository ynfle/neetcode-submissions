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

func countIsValid(count map[byte]int) bool {
	for b, c := range count {
		if c != 1 && b != '.' {
			return false
		}
	}
	return true
}

func buildRowCount(board [][]byte, row int) map[byte]int {
	count := map[byte]int{}
	for _, b := range board[row] {
		count[b] += 1
	}
	return count
}

func buildColCount(board [][]byte, col int) map[byte]int {
	count := map[byte]int{}
	for row := 0; row < 9; row++ {
		b := board[row][col]
		count[b] += 1
	}
	return count
}

func buildBoxCount(board [][]byte, boxRow int, boxCol int) map[byte]int {
	count := map[byte]int{}
	for row := boxRow*3; row < boxRow*3+3; row++ {
		for col := boxCol*3; col < boxCol*3+3; col++ {
			b := board[row][col]
			count[b] += 1
		}
	}
	return count
}