package main

import (
	"os"

	"sudoku"
)

func main() {
	switch {
	case len(os.Args[1:]) != 9:
		sudoku.Mymknch()
		return
	}

	grid := make([][]int, 9)
	for i := range grid {
		grid[i] = make([]int, 9)
	}

	for i, arg := range os.Args[1:] {
		for ii, c := range arg {
			switch {
			case (c < '1' || c > '9') && c != '.':
				sudoku.Mymknch()
				return
			case c >= '1' && c <= '9':
				grid[i][ii] = int(c - 48)
			case c == '.':
				grid[i][ii] = 0
			}
		}
	}

	for column := 0; column < 9; column++ {
		for row1 := 0; row1 < 8; row1++ {
			for row2 := row1 + 1; row2 < 8; row2++ {
				switch {
				case grid[row1][column] == grid[row2][column] && grid[row1][column] > 0:
					sudoku.Mymknch()
					return
				}
			}
		}
	}
	switch {
	case sudoku.Solve(grid):
		sudoku.SudokuTable(grid)
	default:
		sudoku.Mymknch()

	}
}
