package piscine

// SortWordArr function sorts a string array by ascii  in asc order.
func SortWordArr(array []string) {

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
