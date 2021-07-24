package piscine

func MakeRange(min, max int) []int {

	if min < max {

		size := max - min

		numbers := make([]int, size)

		for i := 0; i < size; i++ {
			numbers[i] = i + min
		}
		return numbers

	} else {
		return nil
	}

}
