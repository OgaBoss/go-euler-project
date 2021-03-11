package multiple_3_5

func GetSumOfMultiple35(size int) int {
	sum := 0

	for size > 0 {
		if size%3 == 0 || size%5 == 0 {
			sum = sum + size
		}
		size--
	}

	return sum
}
