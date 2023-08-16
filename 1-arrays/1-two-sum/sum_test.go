package twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ( 
	input       = []int{2, 7, 11, 15}
	target      = 9
	expectedOut = []int{0, 1}
)

func TestTwoSum(t *testing.T) {
	result := twoSum(input, target)
	assert.Equal(t, result, expectedOut, fmt.Sprintf("Expected %d, got %d", expectedOut, result))
}
