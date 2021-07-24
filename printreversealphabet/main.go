package main

import "github.com/01-edu/z01"

func main() {
	i := 25
	var aRune string = "abcdefghijklmnopqrstuvwxyz"
	for i >= 0 {
		z01.PrintRune(rune(aRune[i]))
		i = i - 1
	}
	z01.PrintRune('\n')
}
