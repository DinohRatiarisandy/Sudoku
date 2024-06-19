package main

import "github.com/DinohRatiarisandy/Sudoku/sudoku"

func main() {
	newSudoku := sudoku.Generate()
	newSudoku.Print()
}
