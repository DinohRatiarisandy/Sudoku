package checker

// Can take place at board[i][j] ?
// -> Unique in row, column and slot 3x3
func IsValidToPlace(sudoku *[9][9]int, rowIdx, colIdx, numberToPlace int) bool {
	return isUniqueOnRow(sudoku, rowIdx, numberToPlace) &&
		isUniqueOnColumn(sudoku, colIdx, numberToPlace) &&
		isUniqueOnBox(sudoku, rowIdx, colIdx, numberToPlace)
}

func isUniqueOnRow(sudoku *[9][9]int, rowIdx, numberToPlace int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[rowIdx][i] == numberToPlace {
			return false
		}
	}
	return true
}

func isUniqueOnColumn(sudoku *[9][9]int, colIdx, numberToPlace int) bool {
	for j := 0; j < 9; j++ {
		if sudoku[j][colIdx] == numberToPlace {
			return false
		}
	}
	return true
}

func isUniqueOnBox(sudoku *[9][9]int, rowIdx, colIdx, numberToPlace int) bool {
	var startRow int = rowIdx - rowIdx%3
	var startCol int = colIdx - colIdx%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[startRow+i][startCol+j] == numberToPlace {
				return false
			}
		}
	}
	return true
}
