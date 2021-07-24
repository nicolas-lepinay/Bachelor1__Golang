package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
    result := ""
    for i := 0; i < len(a); i++ {
        result += a[i] + "\n"
    }
    runes := []rune(result)
    for j := range runes {
        z01.PrintRune(runes[j])
    }
}
