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
