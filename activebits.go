package piscine

// ActiveBits function returns the number of active bits
// (bits with the value 1) in the binary representation of an integer.

func ActiveBits(n int) int {
	var active int

	for n != 0 {
		mod := n % 2
		if mod == 1 {
			active++
		}
		n = n / 2
	}

	return active
}
