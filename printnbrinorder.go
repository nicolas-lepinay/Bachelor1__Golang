package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var nombre []int
		valeur := n
		count := 0
		for n != 0 {
			valeur = n % 10
			n /= 10
			nombre = append(nombre, valeur)
			count++
		}
		for i, j := 0, 1; j < count; i, j = i+1, j+1 {
			if nombre[i] > nombre[j] {
				nombre[i], nombre[j] = nombre[j], nombre[i]
				i = -1
				j = 0
			}
		}
		for t := 0; t < count; t++ {
			z01.PrintRune(rune(nombre[t] + 48))
		}
	}
}
