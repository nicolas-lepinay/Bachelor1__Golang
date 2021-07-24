package piscine

func Map(f func(int) bool, a []int) []bool {

	size := len(a)
	result := make([]bool, size)

	for i, num := range a {
		result[i] = f(num)
	}

	return result
}
