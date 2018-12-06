package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	i := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
	o := 17
	assert.Equal(t, o, solve1(i))
}

func TestSolve2(t *testing.T) {
	i := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
	o := 16
	assert.Equal(t, o, solve2(i, 32))
}
