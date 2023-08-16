package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	input = "abccdba"
	expect = true
)

func TestPalindrome(t *testing.T) {
	actual := validPalindrome(input)
	assert.Equal(t, expect, actual)
}
