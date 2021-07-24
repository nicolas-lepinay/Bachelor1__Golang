package piscine

func Any(f func(string) bool, arr []string) bool {

	for _, word := range arr {
		if f(word) {
			return true
		}
	}
	return false
}
