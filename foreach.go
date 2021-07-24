package piscine

func ForEach(f func(int), arr []int) {
	for _, num := range arr {
		f(num)
	}
}
