package piscine

func NRune(s string, n int) rune {
	if len(s) == 0 {
		return 0
	}
	if n > len(s) || n <= 0 {
		return 0
	}
	i := n - 1
	runes := []rune(s)
	return runes[i]
}
