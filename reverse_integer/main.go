package reverse_integer

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	result := 0

	for x != 0 {
		rem := x % 10
		x = x / 10

		result = (result * 10) + rem
	}

	// Max value check
	if result >= 2147483648 || result <= -2147483648 {
		return 0
	}

	return result
}
