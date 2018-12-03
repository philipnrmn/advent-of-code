package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	i := "abcdef bababc abbcde abcccd aabcdd abcdee ababab"
	o := 12
	assert.Equal(t, o, solve1(i))
}

func TestSolve2(t *testing.T) {
	i := "abcde fghij klmno pqrst fguij axcye wvxyz"
	o := "fgij"
	assert.Equal(t, o, solve2(i))
}
