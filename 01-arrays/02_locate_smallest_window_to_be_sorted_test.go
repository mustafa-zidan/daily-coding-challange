package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestSortedWindow(t *testing.T) {
	// case 1 example 1
	input := []int{3, 7, 5, 6, 9}
	start, end := smallestSortedWindow(input)
	assert.Equal(t, 1, start)
	assert.Equal(t, 3, end)
	// case 2 shifted by one
	input = []int{3, 9, 5, 6, 7}
	start, end = smallestSortedWindow(input)
	assert.Equal(t, 1, start)
	assert.Equal(t, 4, end)
	// case 3 fully sorted
	input = []int{1, 2, 3, 4, 5, 6, 7}
	start, end = smallestSortedWindow(input)
	assert.Equal(t, -1, start)
	assert.Equal(t, -1, end)
	// case 4 inversed
	input = []int{7, 6, 5, 4, 3, 2, 1, 0}
	start, end = smallestSortedWindow(input)
	assert.Equal(t, 0, start)
	assert.Equal(t, 8, end)
	// case 5 Empty
	input = make([]int, 0)
	start, end = smallestSortedWindow(input)
	assert.Equal(t, -1, start)
	assert.Equal(t, -1, end)
}
