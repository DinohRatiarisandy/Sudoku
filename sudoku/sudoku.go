package sudoku

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/DinohRatiarisandy/Sudoku/checker"
	"github.com/gin-gonic/gin"
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

func CreateSudokuPuzzle(c *gin.Context) {
	difficulty := c.DefaultQuery("difficulty", "easy")
	newSudokuPuzzle := generateNewSudoku(difficulty)
	c.JSON(http.StatusOK, gin.H{
		"state":      toOneDimension(newSudokuPuzzle.State),
		"solution":   toOneDimension(newSudokuPuzzle.Solution),
		"difficulty": newSudokuPuzzle.Difficulty,
	})
}

func generateNewSudoku(difficulty string) *Sudoku {
	min, max := getDifficulty(difficulty)
	numToRemove := 81 - (rand.Intn(max-min+1) + min)

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

func toOneDimension(array2D [9][9]int) string {
	res := ""

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			res += strconv.Itoa(array2D[i][j])
		}
	}
	return res
}
