package arrays

import (
	"testing"

	"github.com/pkg/profile"
	"github.com/stretchr/testify/assert"
)

func TestProductOfOtherElements(t *testing.T) {
	// case 1 example 1
	defer profile.Start().Stop()
	input := []int{1, 2, 3, 4, 5}
	output := productOfTheOthers(input)
	assert.Equal(t, []int{120, 60, 40, 30, 24}, output)
	// case 2 example 2
	input = []int{3, 2, 1}
	output = productOfTheOthers(input)
	assert.Equal(t, []int{2, 3, 6}, output)
	// case 3 zero element
	input = []int{3, 2, 1, 0}
	output = productOfTheOthers(input)
	assert.Equal(t, []int{2, 3, 6, 6}, output)
	// case 4 two zero elements
	input = []int{0, 3, 2, 1, 0}
	output = productOfTheOthers(input)
	assert.Equal(t, []int{6, 2, 3, 6, 6}, output)
}

func TestProductOfOtherElementsNoDevision(t *testing.T) {
	defer profile.Start().Stop()
	// case 1 example 1
	input := []int{1, 2, 3, 4, 5}
	output := productOfTheOthersNoDivision(input)
	assert.Equal(t, []int{120, 60, 40, 30, 24}, output)
	// case 2 example 2
	input = []int{3, 2, 1}
	output = productOfTheOthersNoDivision(input)
	assert.Equal(t, []int{2, 3, 6}, output)
	// case 3 zero element
	input = []int{3, 2, 1, 0}
	output = productOfTheOthersNoDivision(input)
	assert.Equal(t, []int{2, 3, 6, 6}, output)
	// case 4 two zero elements
	input = []int{0, 3, 2, 1, 0}
	output = productOfTheOthersNoDivision(input)
	assert.Equal(t, []int{6, 2, 3, 6, 6}, output)
}
