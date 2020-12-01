package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	i := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	o := 4
	assert.Equal(t, o, solve1(i))
}

func TestSolve2(t *testing.T) {
	i := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	o := 3
	assert.Equal(t, o, solve2(i))
}
