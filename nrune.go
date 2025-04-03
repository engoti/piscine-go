package piscine

func NRune(s string, n int) rune {
	myrune := []rune(s)
	length := len(s)

	if n > 0 && n <= length {
		return myrune[n-1]
	}
	return 0
}
