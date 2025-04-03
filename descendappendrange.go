package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	q := []int{}
	for i := max; i > min; i-- {
		q = append(q, i)
	}
	return q
}
