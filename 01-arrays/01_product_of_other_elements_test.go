package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfOtherElements(t *testing.T) {
	// case 1 example 1
	input := []int{1, 2, 3, 4, 5}
	output := Product(input)
	assert.Equal(t, []int{120, 60, 40, 30, 24}, output)
	// case 2 example 2
	input = []int{3, 2, 1}
	output = Product(input)
	assert.Equal(t, []int{2, 3, 6}, output)
	// case 3 zero element
	input = []int{3, 2, 1, 0}
	output = Product(input)
	assert.Equal(t, []int{2, 3, 6, 6}, output)
	// case 4 two zero elements
	input = []int{0, 3, 2, 1, 0}
	output = Product(input)
	assert.Equal(t, []int{6, 2, 3, 6, 6}, output)
}
