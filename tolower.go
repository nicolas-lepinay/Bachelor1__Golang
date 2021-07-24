package piscine

func ToLower(s string) string {

	chaine := []rune(s)

	for i := 0; i < len(s); i++ {
		if chaine[i] >= 65 && chaine[i] <= 90 {
			chaine[i] = chaine[i] + 32
		}
	}
	return string(chaine)
}
