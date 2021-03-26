package Euler4

import (
	"fmt"
	"strconv"
)

func isPalindrome(number int) bool {
	strNumber := strconv.Itoa(number)
	reverseStrNumber := ""

	for length := len(strNumber); length > 0; length-- {
		reverseStrNumber += string(strNumber[length-1])
	}
	reverseNum, err := strconv.Atoi(reverseStrNumber)
	if err != nil {
		fmt.Println("Failure to cast String to int")
	}

	return number == reverseNum
}

func GetMaximumPalindrome() int {
	max := 100
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindrome(product) && product > max {
				max = product
			}
		}
	}

	return max
}
