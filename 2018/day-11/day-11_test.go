package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	io := map[string][]int{
		"18": []int{33, 45},
		"42": []int{21, 61},
	}
	for s, coords := range io {
		x, y := solve1(s)
		assert.Equal(t, coords[0], x)
		assert.Equal(t, coords[1], y)
	}
}

func TestLevelOf(t *testing.T) {
	i := [][]int{
		[]int{3, 5, 8},
		[]int{122, 79, 57},
		[]int{217, 196, 39},
		[]int{101, 153, 71},
	}
	o := []int{4, -5, 0, 4}
	for n, args := range i {
		assert.Equal(t, o[n], levelOf(args[0], args[1], args[2]))
	}
}
