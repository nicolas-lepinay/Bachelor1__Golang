package piscine

func IsUpper(s string) bool {

	string := []rune(s)

	for i := 0; i < len(s); i++ {
		if string[i] >= 65 && string[i] <= 90 {

		} else {
			return false
		}
	}
	return true
}
