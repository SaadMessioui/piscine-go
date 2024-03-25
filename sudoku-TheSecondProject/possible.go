package sudoku

func Possible(grid [][]int, row, column, number int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == number {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if grid[i][column] == number {
			return false
		}
	}

	x0 := (column / 3) * 3
	y0 := (row / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[y0+i][x0+j] == number {
				return false
			}
		}
	}
	return true
}
