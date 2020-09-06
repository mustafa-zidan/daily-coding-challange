package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SmallerCountSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSmallerCount(t *testing.T) {
	suite.Run(t, new(SmallerCountSuite))
}

func (suite *SmallerCountSuite) TestSmallerCount() {
	// Example 1
	input := []int{3, 4, 9, 6, 1}
	output := smallerCount(input)
	assert.Equal(suite.T(), []int{1, 1, 2, 1, 0}, output)
}

func (suite *SmallerCountSuite) TestSmallerCountBrute() {
	// Example 1
	input := []int{3, 4, 9, 6, 1}
	output := smallerCountBrute(input)
	assert.Equal(suite.T(), []int{1, 1, 2, 1, 0}, output)
}

func BenchmarkSmallerCount(b *testing.B) {
	input := []int{3, 4, 9, 6, 1}
	for i := 0; i < b.N; i++ {
		smallerCount(input)
	}
}

func BenchmarkSmallerCountBrute(b *testing.B) {
	input := []int{3, 4, 9, 6, 1}
	for i := 0; i < b.N; i++ {
		smallerCountBrute(input)
	}
}
