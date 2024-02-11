package piscine

func Atoi(s string) int {
	result := 0
	sing := 1
	for i, v := range s {
		switch {
		case i == 0 && v == '+':
			sing = 1
		case i == 0 && v == '-':
			sing = -1
		case v >= '0' && v <= '9':
			result = result*10 + int(v-'0')
		default:
			return 0
		}
	}
	return result * sing
}
