package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input = "abcbad"
	expect = 4
)

func TestLongest(t *testing.T) {
	actual := lengthOfLongestSubstring(input)
	assert.Equal(t, expect, actual)
}
