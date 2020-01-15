package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagramIndicies(t *testing.T) {
	w, s, expected := "ab", "abxaba", []int{0, 3, 4}
	output := findAnagramIndicies(w, s)
	assert.Equal(t, expected, output)
}
