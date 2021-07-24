package piscine

func IsNumeric(s string) bool {

	string := []rune(s)

	for i := 0; i < len(s); i++ {
		if string[i] >= 48 && string[i] <= 57 {

		} else {
			return false
		}
	}
	return true
}
