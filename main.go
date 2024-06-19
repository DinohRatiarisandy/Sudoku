package main

import (
	"fmt"

	"github.com/DinohRatiarisandy/Sudoku/tools/generator"
)

func showResult(sudoku *[9][9]int) {
	for _, row := range sudoku {
		fmt.Println(row)
	}
}

func main() {
	var sudoku = generator.GenerateNewSudoku()
	showResult(&sudoku.State)
}
