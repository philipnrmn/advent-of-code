package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	io := map[string]int{
		"+1 -2 +3 +1": 3,
		"+1 +1 +1":    3,
		"+1 +1 -2":    0,
		"-1 -2 -3":    -6,
	}
	for i, o := range io {
		assert.Equal(t, o, solve1(i))
	}
}

func TestSolve2(t *testing.T) {
	io := map[string]int{
		"+1 -2 +3 +1":    2,
		"+1 -1":          0,
		"+3 +3 +4 -2 -4": 10,
		"-6 +3 +8 +5 -6": 5,
		"+7 +7 -2 -7 -4": 14,
	}
	for i, o := range io {
		assert.Equal(t, o, solve2(i))
	}
}
