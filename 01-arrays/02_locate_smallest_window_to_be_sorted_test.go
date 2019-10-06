package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestSortedWindow(t *testing.T) {
	// case 1 example 1
	input := []int{3, 7, 5, 6, 9}
	start, end := smallestSortedWindow(input)
	assert.Equal(t, 1, start)
	assert.Equal(t, 3, end)
}
