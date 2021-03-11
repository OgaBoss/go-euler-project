package Euler2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	assert.Equal(t, Fibonacci(), 4613732)
}
