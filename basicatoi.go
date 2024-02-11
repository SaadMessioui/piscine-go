package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, v := range s {
		switch {
		case v >= '0' && v <= '9':
			result = result*10 + int(v-'0')

		}
	}
	return result
}
