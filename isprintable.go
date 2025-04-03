package piscine

func IsPrintable(s string) bool {
	for _, q := range s {
		if q < ' ' || q > '~' {
			return false
		}
	}
	return true
}
