package piscine

func ConcatParams(args []string) string {
	result := ""

	for i := 0; i < len(args); i++ {
		//Pour tous les mots sauf le dernier, on ajoute un saut à la ligne :
		if i < len(args)-1 {
			result = result + args[i] + "\n"

		} else {
			//Pour le dernier mot, il ne faut pas de saut à la ligne :
			result = result + args[i]
		}
	}
	return result
}
