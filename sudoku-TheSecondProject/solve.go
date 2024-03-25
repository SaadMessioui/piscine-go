package sudoku

func Solve(grid [][]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			switch {
			case grid[row][column] == 0:
				for number := 1; number <= 9; number++ {
					switch {
					case Possible(grid, row, column, number):
						grid[row][column] = number
						if Solve(grid) {
							return true
						}
						grid[row][column] = 0
					}
				}
				return false
			}
		}
	}

	return true
}
