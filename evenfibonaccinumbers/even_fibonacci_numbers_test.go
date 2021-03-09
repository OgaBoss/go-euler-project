package evenfibonaccinumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	assert.Equal(t, Fibonacci(20), 233168)
}
