package piscine

func IsAlpha(s string) bool {
	my_arr := []rune(s)
	my_check := true
	my_len := 0

	for i := range my_arr {
		my_len = i
	}

	for i := 0; i <= my_len; i++ {
		if (my_arr[i] < 'A' || my_arr[i] > 'z') &&
			(my_arr[i] < '0' || my_arr[i] > '9') {
			my_check = false
			break
		}
	}
	return my_check
}
