package piscine

func IsNumeric(s string) bool {
	my_arr := []rune(s)
	ln := 0
	for i := range my_arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if my_arr[i] < '0' || my_arr[i] > '9' {
			return false
		}
	}
	return true
}
