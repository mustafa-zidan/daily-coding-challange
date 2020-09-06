package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MathSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMath(t *testing.T) {
	suite.Run(t, new(MathSuite))
}

func (suite *MathSuite) TestMax() {
	assert.Equal(suite.T(), 2, Max(1, 2))
	assert.Equal(suite.T(), 2, Max(2, 1))
	assert.Equal(suite.T(), 2, Max(2, 2))
}

func (suite *MathSuite) TestMin() {
	assert.Equal(suite.T(), 1, Min(1, 2))
	assert.Equal(suite.T(), 1, Min(2, 1))
	assert.Equal(suite.T(), 1, Min(1, 1))
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(1, 2)
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(1, 2)
	}
}
