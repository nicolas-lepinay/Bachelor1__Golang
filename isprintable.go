package piscine

func IsPrintable(s string) bool {

	string := []rune(s)

	for i := 0; i < len(s); i++ {
		if string[i] >= 32 && string[i] <= 126 {

		} else {
			return false
		}
	}
	return true
}
