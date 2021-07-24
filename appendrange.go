package piscine

func AppendRange(min, max int) []int {

	var numbers []int

	for i := min; i < max; i++ {
		numbers = append(numbers, i)
	}
	return numbers

}
