package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	i := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
	o := 138
	assert.Equal(t, o, solve1(i))
}

func TestSolve2(t *testing.T) {
	i := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
	o := 66
	assert.Equal(t, o, solve2(i))
}
