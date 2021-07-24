package piscine

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
    if len(args) != 3 {
		return
	}

	result := 0
	for _, char := range args[0] {
		if !(rune(char) >= '0' && rune(char) <= '9' || rune(char) == '-') {
			fmt.Println(result)
			return
		}
	}
	for _, char := range args[2] {
		if !(rune(char) >= '0' && rune(char) <= '9') || rune(char) == '-' {
			fmt.Println(result)
			return
		}
	}

	if args[1] == "+" {
		result = Atoi(args[0]) + Atoi(args[2])
	}
	if args[1] == "-" {
		result = Atoi(args[0]) - Atoi(args[2])
	}
	if args[1] == "*" {
		result = Atoi(args[0]) * Atoi(args[2])
	}
	if args[1] == "/" {
		if args[2] == "0" {
			fmt.Println("No division by 0")
			return
		}
		result = Atoi(args[0]) / Atoi(args[2])
	}
	if args[1] == "%" {
		if args[2] == "0" {
			fmt.Println("No Modulo by 0")
			return
		}
		result = Atoi(args[0]) % Atoi(args[2])
	}

	fmt.Println(result)

}

func Atoi(s string) int {

	result := 0

	for _, character := range s {
		result = result*10 + (int(character - 48))
	}
	return result
}
