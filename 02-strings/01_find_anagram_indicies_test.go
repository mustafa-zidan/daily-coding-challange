package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagramIndicies(t *testing.T) {
	w, s, expected := "ab", "abxaba", []int{0, 3, 4}
	output := findAnagramIndicies(s, w)
	assert.Equal(t, expected, output)
}
