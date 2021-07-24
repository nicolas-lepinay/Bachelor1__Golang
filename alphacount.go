package piscine

func AlphaCount(s string) int {

	lettre := []rune(s)
	count := 0
	for n := 0; n < len(s); n++ {
		if lettre[n] >= 65 && lettre[n] <= 90 || lettre[n] >= 97 && lettre[n] <= 122 {
			count++
		}
	}
	return count
}
