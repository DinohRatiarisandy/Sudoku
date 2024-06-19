package sudoku

import (
	"fmt"
	"math/rand"

	"github.com/DinohRatiarisandy/Sudoku/checker"
)

type Sudoku struct {
	State      [9][9]int
	Solution   [9][9]int
	Difficulty string
}

func getDifficulty(difficulty string) (min, max int) {
	switch difficulty {
	case "easy":
		min = 36
		max = 49
	case "medium":
		min = 32
		max = 35
	case "hard":
		min = 28
		max = 31
	case "evil":
		min = 17
		max = 27
	}
	return min, max
}

func GenerateNewSudoku(difficulty string) *Sudoku {
	min, max := getDifficulty(difficulty)
	numToRemove := 81 - (rand.Intn(max-min+1) + min)
	fmt.Println("Number removed:", numToRemove)
	fmt.Println("Difficulty:", difficulty)

	var newSudoku Sudoku
	newSudoku.solve()
	newSudoku.Solution = newSudoku.State
	newSudoku.Difficulty = difficulty

	idxCellsToRemove := rand.Perm(81)

	for _, cell := range idxCellsToRemove {
		row := cell / 9
		col := cell % 9
		backup := newSudoku.State[row][col]
		newSudoku.State[row][col] = 0

		if newSudoku.countSolution() != 1 {
			newSudoku.State[row][col] = backup
		} else {
			numToRemove -= 1
			if numToRemove <= 0 {
				break
			}
		}
	}
	return &newSudoku
}

func (sudoku *Sudoku) Print(what string) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if what == "solution" {
				fmt.Print(sudoku.Solution[row][col], "  ")
			} else {
				if sudoku.State[row][col] > 0 {
					fmt.Print(sudoku.State[row][col], "  ")
				} else {
					fmt.Print("   ")
				}
			}
			if col < 8 && col%3 == 2 {
				fmt.Print("| ")
			}
		}
		fmt.Print("\n")
		if row < 8 && row%3 == 2 {
			fmt.Println("---------+----------+---------")
		}
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
