package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MaxSubarraySumSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMaxSubarraySum(t *testing.T) {
	suite.Run(t, new(MaxSubarraySumSuite))
}

func (suite *SmallestSortedWindowsSuite) TestMaxSubarraySum() {
	// Example 1
	input := []int{34, -50, 42, 14, -5, 86}
	output := maxSubarraySum(input)
	assert.Equal(suite.T(), 137, output)

	// Example 2
	input = []int{-5, -1, -8, -9}
	output = maxSubarraySum(input)
	assert.Equal(suite.T(), 0, output)
}

func (suite *SmallestSortedWindowsSuite) TestMaxCircularSubarraySum() {
	// Example 1
	input := []int{34, -50, 42, 14, -5, 86}
	output := maxCircularSubarraySum(input)
	assert.Equal(suite.T(), 171, output)

	// Example 2
	input = []int{-5, -1, -8, -9}
	output = maxCircularSubarraySum(input)
	assert.Equal(suite.T(), 0, output)

	// Example 3
	input = []int{8, -1, 3, 4}
	output = maxCircularSubarraySum(input)
	assert.Equal(suite.T(), 15, output)
}

func BenchmarkMaxSubarraySum(b *testing.B) {
	input := []int{34, -50, 42, 14, -5, 86}
	for i := 0; i < b.N; i++ {
		maxSubarraySum(input)
	}
}

func BenchmarkMaxCircularSubarraySum(b *testing.B) {
	input := []int{34, -50, 42, 14, -5, 86}
	for i := 0; i < b.N; i++ {
		maxCircularSubarraySum(input)
	}
}
