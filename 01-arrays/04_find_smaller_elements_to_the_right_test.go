package arrays

import (
	"testing"

	"github.com/pkg/profile"
	"github.com/stretchr/testify/assert"
)

func TestSmallerCount(t *testing.T) {
	// Example 1
	defer profile.Start().Stop()
	input := []int{3, 4, 9, 6, 1}
	output := smallerCount(input)
	assert.Equal(t, []int{1, 1, 2, 1, 0}, output)
}

func TestSmallerCountBrute(t *testing.T) {
	// Example 1
	defer profile.Start().Stop()
	input := []int{3, 4, 9, 6, 1}
	output := smallerCountBrute(input)
	assert.Equal(t, []int{1, 1, 2, 1, 0}, output)
}
