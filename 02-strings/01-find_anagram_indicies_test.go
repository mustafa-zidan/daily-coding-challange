package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FindAnagramIndiciesSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFindAnagramIndicies(t *testing.T) {
	suite.Run(t, new(FindAnagramIndiciesSuite))
}

func (suite *FindAnagramIndiciesSuite) TestFindAnagramIndicies() {
	w, s, expected := "ab", "abxaba", []int{0, 3, 4}
	output := findAnagramIndicies(s, w)
	assert.Equal(suite.T(), expected, output)
}

func BenchmarkFindAnagramIndicies(b *testing.B) {
	w, s := "ab", "abxaba"
	for i := 0; i < b.N; i++ {
		findAnagramIndicies(s, w)
	}
}
