package Euler3

func primeFactors(number int) []int {
	factors := make([]int, 0)

	for i := 2; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
			number /= i
			i--
		}
	}

	return factors
}

func findLargest(factors []int) int {
	max := factors[0]

	for i := 1; i < len(factors); i++ {
		if factors[i] > factors[i-1] {
			max = factors[i]
		}
	}

	return max
}

func FindPrimeFactors(number int) int {
	return findLargest(primeFactors(number))
}
