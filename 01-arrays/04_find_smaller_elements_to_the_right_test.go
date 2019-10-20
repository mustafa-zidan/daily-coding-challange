package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallerCount(t *testing.T) {
	// Example 1
	input := []int{3, 4, 9, 6, 1}
	output := smallerCount(input)
	assert.Equal(t, []int{1, 1, 2, 1, 0}, output)
}
