package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SmallestCountSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSmallestCount(t *testing.T) {
	suite.Run(t, new(SmallestCountSuite))
}

func (suite *SmallestCountSuite) TestSmallestCount() {
	// Example 1
	input := []int{3, 4, 9, 6, 1}
	output := smallestCount(input)
	assert.Equal(suite.T(), []int{1, 1, 2, 1, 0}, output)
}

func (suite *SmallestCountSuite) TestSmallestCountBrute() {
	// Example 1
	input := []int{3, 4, 9, 6, 1}
	output := smallestCountBrute(input)
	assert.Equal(suite.T(), []int{1, 1, 2, 1, 0}, output)
}

func BenchmarkSmallestCount(b *testing.B) {
	input := []int{3, 4, 9, 6, 1}
	for i := 0; i < b.N; i++ {
		smallestCount(input)
	}
}

func BenchmarkSmallestCountBrute(b *testing.B) {
	input := []int{3, 4, 9, 6, 1}
	for i := 0; i < b.N; i++ {
		smallestCountBrute(input)
	}
}
