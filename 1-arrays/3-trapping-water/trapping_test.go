package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var (
// 	input  = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
// 	expect = 6
// )

func TestTrapping (t *testing.T) {
	actual := trap(input)

	assert.Equal(t, expect, actual)
}
