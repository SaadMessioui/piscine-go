package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, v := range s {

		result = result*10 + int(v-'0')

	}

	return result
}
