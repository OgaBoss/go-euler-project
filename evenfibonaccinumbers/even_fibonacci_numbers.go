package evenfibonaccinumbers

func Fibonacci() int {
	total := 0
	fibTotal := 0
	var cache []int

	for term := 0; fibTotal <= 4000000; term++ {
		if term == 0 || term == 1 {
			cache = append(cache, 1)
		} else {
			fib := cache[term-1] + cache[term-2]
			fibTotal = fib
			cache = append(cache, fib)

			if fib%2 == 0 {
				total = total + fib
			}
		}
	}

	return total
}
