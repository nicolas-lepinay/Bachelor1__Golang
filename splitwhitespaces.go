package piscine

func SplitWhiteSpaces(s string) []string {

	lettres := ""
	var tableau []string

	for i, char := range s {
		if char != ' ' && char != '\t' && char != '\n' {
			lettres += string(s[i])
		} else if lettres != "" {
			tableau = append(tableau, lettres)
			lettres = ""
		}
	}
	tableau = append(tableau, lettres)
	return tableau
}
