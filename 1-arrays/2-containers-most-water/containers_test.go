package containersmostwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input  = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expect = 49
)

func TestContainers(t *testing.T) {
	actual := maxArea(input)
	assert.Equal(t, expect, actual, "test failed")
}
