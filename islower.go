package piscine

func IsLower(s string) bool {

	string := []rune(s)

	for i := 0; i < len(s); i++ {
		if string[i] >= 97 && string[i] <= 122 {

		} else {
			return false
		}
	}
	return true
}
