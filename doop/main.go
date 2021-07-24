package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]

	if len(argument) < 3 || len(argument) > 4 {
		os.Exit(0)
	}
	var1 := ToInt(os.Args[1])
	signe := os.Args[2]
	var2 := ToInt(os.Args[3])
	result := 0

	if signe == "*" {
		result = var1 * var2
	} else if signe == "-" {
		result = var1 - var2
	} else if signe == "+" {
		result = var1 + var2
	} else if signe == "/" {
		if var2 == 0 {
			strdiv := "No division by 0"
			for ii := 0; ii < len(strdiv); ii++ {
				z01.PrintRune(rune(strdiv[ii]))
			}
			z01.PrintRune(10)
			os.Exit(0)
		}
		result = var1 / var2
	} else if signe == "%" {
		if int(var2) == 0 {
			strmod := "No modulo by 0"
			for ii := 0; ii < len(strmod); ii++ {
				z01.PrintRune(rune(strmod[ii]))
			}
			z01.PrintRune(10)
			os.Exit(0)
		}
		result = var1 % var2
	} else {
		os.Exit(0)
	}
	if result > 2147483647 || result < -2147483648 {
		os.Exit(0)
	}
	var res []int
	n := result
	count := 0
	j := 10

	for n != 0 {
		j = n % 10
		n /= 10
		if result > 0 {
			res = append(res, j)
		}
		if result < 0 {
			res = append(res, -j)
		}
		count++
	}
	for z := count - 1; z >= 0; z-- {
		if result > 0 {
			z01.PrintRune(rune(res[z] + 48))
		}
		if result < 0 {
			if z == count-1 {
				z01.PrintRune(45)
			}
			z01.PrintRune(rune(res[z] + 48))
		}
	}
	if result == 0 {
		z01.PrintRune(48)
	}
	z01.PrintRune(10)
}

func ToInt(a string) int {
	var b int
	mult := 1
	inv := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] >= 48 && a[i] <= 57 || a[i] == 45 {
		} else {
			os.Exit(0)
		}
		if a[i] == 45 {
			inv = 1
		} else {
			b += (int(a[i]) - 48) * mult
			mult *= 10
		}
	}
	if inv == 1 {
		b = -b
	}
	return b
}
