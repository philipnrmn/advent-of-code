package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	io := map[string]int{
		"1122":     3,
		"1111":     4,
		"1234":     0,
		"91212129": 9,
	}
	for i, o := range io {
		assert.Equal(t, solve1(i), o)
	}
}

func TestSolve2(t *testing.T) {
	io := map[string]int{
		"1212":     6,
		"1221":     0,
		"123425":   4,
		"123123":   12,
		"12131415": 4,
	}
	for i, o := range io {
		assert.Equal(t, solve2(i), o)
	}
}
