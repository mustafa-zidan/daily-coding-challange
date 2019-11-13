package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarraySum(t *testing.T) {
	// Example 1
	input := []int{34, -50, 42, 14, -5, 86}
	output := maxSubarraySum(input)
	assert.Equal(t, 137, output)

	// Example 2
	input = []int{-5, -1, -8, -9}
	output = maxSubarraySum(input)
	assert.Equal(t, 0, output)
}

func TestMaxCircularSubarraySum(t *testing.T) {
	// Example 1
	input := []int{34, -50, 42, 14, -5, 86}
	output := maxCircularSubarraySum(input)
	assert.Equal(t, 171, output)

	// Example 2
	input = []int{-5, -1, -8, -9}
	output = maxCircularSubarraySum(input)
	assert.Equal(t, 0, output)

	// Example 3
	input = []int{8, -1, 3, 4}
	output = maxCircularSubarraySum(input)
	assert.Equal(t, 15, output)
}
