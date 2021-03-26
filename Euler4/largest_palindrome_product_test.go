package Euler4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestPrimeFactors(t *testing.T) {
	assert.Equal(t, GetMaximumPalindrome(), 6857)
}
