package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	inputS = "ab##"
	inputT = "dc#d#"
	expect = false
)

func TestTwoStrings(t *testing.T) {
	actual := backspaceCompare(inputS, inputT)
	assert.Equal(t, expect, actual)
}
