package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDietPlan(t *testing.T) {
	// Input: nums = [1,2,3,4,5], k = 1, lower = 3, upper = 3
	// Output: 0
	input := []int{1, 2, 3, 4, 5}
	output := DietPlan(input, 1, 3, 3)
	assert.Equal(t, 0, output)
	// Input: nums = [3,2], k = 2, lower = 0, upper = 1
	// Output: 1
	input = []int{2, 3}
	output = DietPlan(input, 2, 0, 1)
	assert.Equal(t, 1, output)
	// Input: nums = [6,5,0,0], k = 2, lower = 1, upper = 5
	// Output: 0
	input = []int{6, 5, 0, 0}
	output = DietPlan(input, 2, 1, 5)
	assert.Equal(t, 0, output)
}
