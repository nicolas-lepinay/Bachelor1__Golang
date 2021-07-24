package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myProgramName := os.Args[0]
	IndiceDuSlash := -1

	for i, valeur := range myProgramName {
		if valeur == '/' {
			IndiceDuSlash = i
		}
	}
	for _, lettre := range myProgramName[IndiceDuSlash+1:] {
		z01.PrintRune(lettre)
	}
	z01.PrintRune('\n')
}
