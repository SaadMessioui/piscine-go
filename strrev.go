package piscine

func StrRev(s string) string {
	result := ""
	for i := StrLen(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
