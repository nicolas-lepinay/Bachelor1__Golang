package piscine

// IsSorted function returns true if the slice of int is sorted,
// and returns false otherwise.
func IsSorted(f func(a, b int) int, tab []int) bool {
	result := true

	lentab := 0
	for i := range tab {
		lentab = i + 1
	}

	for i := 0; i < lentab-2; i++ {
		if f(tab[i], tab[i+1]) >= 0 {
			if f(tab[i+1], tab[i+2]) < 0 {
				result = false
				break
			}
		} else {
			if f(tab[i+1], tab[i+2]) > 0 {
				result = false
				break
			}
		}
	}
	return result
}
