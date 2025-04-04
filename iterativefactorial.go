package piscine

func IterativeFactorial(nb int) int {
	if nb > 21 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	b := 1
	for a := nb; a >= 1; a-- {
		b *= a
	}
	return b
}
