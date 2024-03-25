package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {

	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0:
				switch {
				case j == 0 || j == x-1:
					z01.PrintRune('A')
				default:
					z01.PrintRune('B')
				}
			case i == y-1:
				switch {
				case j == 0 || j == x-1:
					z01.PrintRune('C')
				default:
					z01.PrintRune('B')
				}

			default:
				switch {
				case j == 0 || j == x-1:
					z01.PrintRune('B')
				default:
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune(10)
	}
}
