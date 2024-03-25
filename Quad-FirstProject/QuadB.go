package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {

	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0:
				switch {
				case j == 0:
					z01.PrintRune('/')
				case j == x-1:
					z01.PrintRune('\\')
				default:
					z01.PrintRune('*')
				}
			case i == y-1:
				switch {
				case j == 0:
					z01.PrintRune('\\')
				case j == x-1:
					z01.PrintRune('/')
				default:
					z01.PrintRune('*')
				}

			default:
				switch {
				case j == 0 || j == x-1:
					z01.PrintRune('*')
				default:
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune(10)
	}
}
