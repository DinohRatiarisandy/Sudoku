package main

import (
	"github.com/DinohRatiarisandy/Sudoku/sudoku"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/newsudoku", sudoku.CreateSudokuPuzzle)
	r.Run()
}
