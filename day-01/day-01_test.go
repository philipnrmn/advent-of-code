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
		assert.Equal(t, solve1(i), o)
	}
}
