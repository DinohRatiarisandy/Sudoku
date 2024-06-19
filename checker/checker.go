package checker

// Can take place at board[i][j] ?
// -> Unique in row, column and box 3x3
func IsValidToPlace(sudoku *[9][9]int, rowIdx, colIdx, numberToPlace int) bool {
	return isUniqueInRow(sudoku, rowIdx, numberToPlace) &&
		isUniqueInColumn(sudoku, colIdx, numberToPlace) &&
		isUniqueInBox(sudoku, rowIdx, colIdx, numberToPlace)
}

func isUniqueInRow(sudoku *[9][9]int, rowIdx, numberToPlace int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[rowIdx][i] == numberToPlace {
			return false
		}
	}
	return true
}

func isUniqueInColumn(sudoku *[9][9]int, colIdx, numberToPlace int) bool {
	for j := 0; j < 9; j++ {
		if sudoku[j][colIdx] == numberToPlace {
			return false
		}
	}
	return true
}

func isUniqueInBox(sudoku *[9][9]int, rowIdx, colIdx, numberToPlace int) bool {
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
