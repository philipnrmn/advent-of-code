package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	io := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}
	for i, o := range io {
		assert.Equal(t, isValid(i), o)
	}
}

func TestSolve1(t *testing.T) {
	input := `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`
	assert.Equal(t, solve1(input), 2)
}
