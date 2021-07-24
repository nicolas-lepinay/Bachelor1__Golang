package piscine

func Compact(ptr *[]string) int {

	length := 0

	for _, element := range *ptr {
		if element != "" {
			length++
		}
	}

	i := 0
	compact := make([]string, length)
	for _, element := range *ptr {
		if element != "" {
			compact[i] = element
			i++
		}
	}

	*ptr = compact
	return length
}
