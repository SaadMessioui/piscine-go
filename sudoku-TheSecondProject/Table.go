package sudoku

import "github.com/01-edu/z01"

func SudokuTable(grid [][]int) {
	for _, row := range grid {
		for i, num := range row {
			switch {
			case num == 0:
				z01.PrintRune('.')
			default:
				z01.PrintRune(rune('0' + num))
			}
			if i < len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
