package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args[1:]
	isflag := false
	length := len(arguments)

	if length >= 1 {

		for _, arg := range arguments {
			if arg == "--upper" {
				isflag = true
				arguments = os.Args[2:]
			}
		}

		for _, arg := range arguments {

			nombre := 0

			for _, chiffre := range arg {
				nombre = nombre*10 + int(chiffre-'0')
			}

			if nombre >= 1 && nombre <= 26 {
				if isflag == false {
					z01.PrintRune(rune(nombre + 96))

				} else {
					z01.PrintRune(rune(nombre + 64))
				}

			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
