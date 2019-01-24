package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := `|
v
|
|
|
^
|`

	x, y := solve1(input)
	assert.Equal(t, 0, x)
	assert.Equal(t, 3, y)
}
