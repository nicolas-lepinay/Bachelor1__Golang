package piscine

func NRune(s string, n int) rune {

	lettre := []rune(s)
	if n > 0 && n <= len(s) {
		return lettre[n-1]
	} else {
		return 0
	}
}
