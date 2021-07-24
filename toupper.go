package piscine

func ToUpper(s string) string {

	chaine := []rune(s)

	for i := 0; i < len(s); i++ {
		if chaine[i] >= 97 && chaine[i] <= 122 {
			chaine[i] = chaine[i] - 32
		}
	}
	return string(chaine)
}
