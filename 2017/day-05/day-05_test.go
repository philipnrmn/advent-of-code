package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStep1(t *testing.T) {
	input := []int{0, 3, 0, 1, -3}
	cursors := []int{
		0,
		0,
		1,
		4,
		1,
		5,
	}
	steps := [][]int{
		{0, 3, 0, 1, -3},
		{1, 3, 0, 1, -3},
		{2, 3, 0, 1, -3},
		{2, 4, 0, 1, -3},
		{2, 4, 0, 1, -2},
		{2, 5, 0, 1, -2},
	}

	for i := 1; i < len(steps); i++ {
		cursor, done := step1(input, cursors[i-1])
		assert.Equal(t, steps[i], input)
		assert.Equal(t, cursors[i], cursor)
		assert.Equal(t, i == 5, done)
	}
}

func TestSolve1(t *testing.T) {
	input := `0
3
0
1
-3`
	assert.Equal(t, solve1(input), 5)
}

func TestSolve2(t *testing.T) {
	input := `0
3
0
1
-3`
	assert.Equal(t, solve2(input), 10)
}
