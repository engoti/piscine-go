package piscine

func Map(f func(int) bool, a []int) []bool {
	fin := make([]bool, len(a))

	for i, res := range a {
		fin[i] = f(res)
	}
	return fin
}
