package Euler3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestPrimeFactors(t *testing.T) {
	assert.Equal(t, FindPrimeFactors(600851475143), 6857)
}
