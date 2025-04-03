package piscine

func IsUpper(s string) bool {
	for _, q := range s {
		if q < 'A' || q > 'Z' {
			return false
		}
	}
	return true
}
