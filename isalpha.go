package piscine

func IsAlpha(s string) bool {

	string := []rune(s)

	for i := 0; i < len(s); i++ {
		if string[i] >= 48 && string[i] <= 57 || string[i] >= 65 && string[i] <= 90 || string[i] >= 97 && string[i] <= 122 {

		} else {
			return false
		}
	}
	return true
}
