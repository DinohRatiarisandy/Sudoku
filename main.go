package main

import (
	"fmt"

	"github.com/DinohRatiarisandy/Sudoku/sudoku"
)

func main() {
	difficulty := "medium" // difficulty: easy, medium, hard, evil
	newSudoku := sudoku.GenerateNewSudoku(difficulty)
	newSudoku.Print("state")
	fmt.Println("\nSOLUTION")
	newSudoku.Print("solution")
}
