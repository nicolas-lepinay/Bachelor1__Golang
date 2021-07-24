package piscine

func Abort(a, b, c, d, e int) int {

	// On crée un tableau contenant a, b, c, d et e :
	var tableau []int
	tableau = append(tableau, a, b, c, d, e)

	// On trie le tableau :
	for i := 0; i < len(tableau); i++ {
		for j := i + 1; j < len(tableau); j++ {
			if tableau[i] > tableau[j] {
				tableau[i], tableau[j] = tableau[j], tableau[i]
			}
		}
	}

	// On renvoit la valeur médiane, c'est-à-dire la 3ème (indice 2) :
	return tableau[2]
}
