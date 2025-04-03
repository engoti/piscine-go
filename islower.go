package piscine

func IsLower(s string) bool {
	for _, q := range s {
		if q < 'a' || q > 'z' {
			return false
		}
	}
	return true
}
