package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := []string{"0", "2", "7", "0"}
	assert.Equal(t, solve1(input), 5)
}

func TestStep(t *testing.T) {
	input := []int{0, 2, 7, 0}
	steps := [][]int{
		[]int{0, 2, 7, 0},
		[]int{2, 4, 1, 2},
		[]int{3, 1, 2, 3},
		[]int{0, 2, 3, 4},
		[]int{1, 3, 4, 1},
		[]int{2, 4, 1, 2},
	}

	for i := 1; i < len(steps); i++ {
		step(input)
		assert.Equal(t, steps[i], input)
	}
}
