package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	io := map[int]int{
		1:    0,
		12:   3,
		23:   2,
		1024: 31,
	}
	for i, o := range io {
		assert.Equal(t, o, solve1(i))
	}
}
