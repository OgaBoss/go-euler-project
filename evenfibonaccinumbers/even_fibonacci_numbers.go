package evenfibonaccinumbers

func Fibonacci(term int) int {
	// var cache []int
	if term == 1 || term == 0 {
		// cache = append(cache, 1)
		return term
	}
	return Fibonacci(term-1) + Fibonacci(term-2)
}
