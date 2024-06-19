package generator

import (
	"math/rand"

	"github.com/DinohRatiarisandy/Sudoku/board"
	"github.com/DinohRatiarisandy/Sudoku/tools/checker"
)

type typeBoard = board.Board

func GenerateNewSudoku() typeBoard {
	var newSudoku typeBoard
	solveSudoku(&newSudoku.State)
	return newSudoku
}

func solveSudoku(sudoku *[9][9]int) bool {
	row, col, found := findEmptyCell(sudoku)
	if !found {
		return true
	}

	randNum := rand.Perm(9)
	for _, num := range randNum {
		num += 1
		if checker.IsValidToPlace(sudoku, row, col, num) {
			sudoku[row][col] = num
			if solveSudoku(sudoku) {
				return true
			}
			sudoku[row][col] = 0
		}
	}
	return false
}

func findEmptyCell(sudoku *[9][9]int) (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if sudoku[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}
