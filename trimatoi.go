package piscine

func TrimAtoi(s string) int {
    runes := []rune(s)
    value := 0
    count := 0
    nombre := 0
    mt := 1
    var a []int
    for l := 0; l < len(runes); l++ {
        if 47 < s[l] && s[l] < 58 || s[l] == 43 || s[l] == 45 {
            value = int(s[l])
            a = append(a, value)
            count++
        }
    }
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    for t := 0; t < count; t++ {
        if 47 < a[t] && a[t] < 58 {
            nombre = nombre + (a[t]-48)*mt
            mt *= 10
        }
    }
    if count != 0 && a[(count-1)] == 45 {
        nombre = -nombre
    }
    return (nombre)
}
