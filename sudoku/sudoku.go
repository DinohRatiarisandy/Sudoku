package sudoku

import (
	"fmt"
	"math/rand"

	"github.com/DinohRatiarisandy/Sudoku/checker"
)

type Sudoku struct {
	State [9][9]int
}

func Generate() *Sudoku {
	var newSudoku Sudoku
	newSudoku.solve()
	return &newSudoku
}

func (sudoku *Sudoku) Print() {
	for _, row := range sudoku.State {
		fmt.Println(row)
	}
}

func (sudoku *Sudoku) findEmptyCell() (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if sudoku.State[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

func (sudoku *Sudoku) countSolution() int {
	row, col, found := sudoku.findEmptyCell()
	if !found {
		return 1
	}

	count := 0
	for num := 1; num <= 9; num++ {
		if checker.IsValidToPlace(&sudoku.State, row, col, num) {
			sudoku.State[row][col] = num
			count += sudoku.countSolution()
			sudoku.State[row][col] = 0
			if count > 1 {
				break
			}
		}
	}
	return count
}

func (sudoku *Sudoku) solve() bool {
	row, col, found := sudoku.findEmptyCell()
	if !found {
		return true
	}

	randNum := rand.Perm(9)
	for _, num := range randNum {
		num += 1
		if checker.IsValidToPlace(&sudoku.State, row, col, num) {
			sudoku.State[row][col] = num
			if sudoku.solve() {
				return true
			}
			sudoku.State[row][col] = 0
		}
	}
	return false
}
