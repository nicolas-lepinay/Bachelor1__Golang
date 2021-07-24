package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for k := 0; k < len(args); k++ {
		for _, letter := range args[k] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
