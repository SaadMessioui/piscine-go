package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for x := j + 1; x <= '9'; x++ {
				if i < j && j < x {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(x)
					if i == '7' && j == '8' && x == '9' {
						z01.PrintRune(10)
						continue
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
