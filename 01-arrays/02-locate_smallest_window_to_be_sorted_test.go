package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SmallestSortedWindowsSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSmallestSortedWindow(t *testing.T) {
	suite.Run(t, new(SmallestSortedWindowsSuite))
}

func (suite *SmallestSortedWindowsSuite) TestSmallestSortedWindow() {
	// case 1 example 1
	input := []int{3, 7, 5, 6, 9}
	start, end := smallestSortedWindow(input)
	assert.Equal(suite.T(), 1, start)
	assert.Equal(suite.T(), 3, end)
	// case 2 shifted by one
	input = []int{3, 9, 5, 6, 7}
	start, end = smallestSortedWindow(input)
	assert.Equal(suite.T(), 1, start)
	assert.Equal(suite.T(), 4, end)
	// case 3 fully sorted
	input = []int{1, 2, 3, 4, 5, 6, 7}
	start, end = smallestSortedWindow(input)
	assert.Equal(suite.T(), -1, start)
	assert.Equal(suite.T(), -1, end)
	// case 4 reversed
	input = []int{7, 6, 5, 4, 3, 2, 1, 0}
	start, end = smallestSortedWindow(input)
	assert.Equal(suite.T(), 0, start)
	assert.Equal(suite.T(), 7, end)
	// case 5 Empty
	input = make([]int, 0)
	start, end = smallestSortedWindow(input)
	assert.Equal(suite.T(), -1, start)
	assert.Equal(suite.T(), -1, end)
}

func BenchmarkSmallestSortedWindow(b *testing.B) {
	input := []int{3, 7, 5, 6, 9}
	for i := 0; i < b.N; i++ {
		smallestSortedWindow(input)
	}
}
