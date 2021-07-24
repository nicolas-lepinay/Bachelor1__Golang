package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]

	x := 0

	if len(args) < 1 {
		bytes, err := ioutil.ReadAll(os.Stdin) // On récupère la saisie de l'utilisateur (Stdin) qu'on stocke dans un tableau de bytes (c-a-d un string!)
		if err != nil {
			return
		}
		for range bytes { // Pour chaque byte composant la saisie de l'utilisateur (Stdin : standard input), afficher ce byte (afficher le byte indice x, puis faire x++, etc.)
			z01.PrintRune(rune(bytes[x]))
			x++
		}

		return
	}

	var str string

	for _, arg := range args { // Pour chaque arg contenu dans args,

		content, err := ioutil.ReadFile(arg) // afficher le contenu des fichiers passés en arguments
		if err != nil {                      // S'il y a une erreur (aucun argument, argument inexistant, etc.)
			str += "ERROR: open " + arg + ": no such file or directory\n" // Par exemple, si je tape "go run main.go gibberish" >>> ERROR: open gibberish: No such file or directory
			for _, char := range str {                                    // Pour chaque caractère (rune) de la string str...
				z01.PrintRune(rune(char)) //... afficher ce caractère/rune.
			}
			os.Exit(1) // Chaque run renvoit un statut (en général, 0 si tout s'est bien passé, ou 1 s'il y a eu une erreur). Ici, on fait en sorte qu'il affiche 1 s'il y a une erreur.

		} else {
			for _, char := range content { // Sinon, afficher chaque caractère contenu dans le fichier texte rentré en argument.
				z01.PrintRune(rune(char))
			}
		}
	}

}
