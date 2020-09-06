package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductOfOtherElementsSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestProductOfOtherElements(t *testing.T) {
	suite.Run(t, new(ProductOfOtherElementsSuite))
}

func (suite *ProductOfOtherElementsSuite) TestProductOfOtherElements() {
	// case 1 example 1
	input := []int{1, 2, 3, 4, 5}
	output := productOfTheOthers(input)
	assert.Equal(suite.T(), []int{120, 60, 40, 30, 24}, output)
	// case 2 example 2
	input = []int{3, 2, 1}
	output = productOfTheOthers(input)
	assert.Equal(suite.T(), []int{2, 3, 6}, output)
	// case 3 zero element
	input = []int{3, 2, 1, 0}
	output = productOfTheOthers(input)
	assert.Equal(suite.T(), []int{2, 3, 6, 6}, output)
	// case 4 two zero elements
	input = []int{0, 3, 2, 1, 0}
	output = productOfTheOthers(input)
	assert.Equal(suite.T(), []int{6, 2, 3, 6, 6}, output)
}

func (suite *ProductOfOtherElementsSuite) TestProductOfOtherElementsNoDevision() {
	// case 1 example 1
	input := []int{1, 2, 3, 4, 5}
	output := productOfTheOthersNoDivision(input)
	assert.Equal(suite.T(), []int{120, 60, 40, 30, 24}, output)
	// case 2 example 2
	input = []int{3, 2, 1}
	output = productOfTheOthersNoDivision(input)
	assert.Equal(suite.T(), []int{2, 3, 6}, output)
	// case 3 zero element
	input = []int{3, 2, 1, 0}
	output = productOfTheOthersNoDivision(input)
	assert.Equal(suite.T(), []int{2, 3, 6, 6}, output)
	// case 4 two zero elements
	input = []int{0, 3, 2, 1, 0}
	output = productOfTheOthersNoDivision(input)
	assert.Equal(suite.T(), []int{6, 2, 3, 6, 6}, output)
}

func BenchmarkProductOfOtherElements(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		productOfTheOthers(input)
	}
}
func BenchmarkProductOfOtherElementsNoDevision(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		productOfTheOthersNoDivision(input)
	}
}
