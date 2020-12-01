package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid1(t *testing.T) {
	io := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}
	for i, o := range io {
		assert.Equal(t, isValid1(i), o)
	}
}

func TestSolve1(t *testing.T) {
	input := `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`
	assert.Equal(t, solve1(input), 2)
}

func TestIsValid2(t *testing.T) {
	io := map[string]bool{
		"abcde fghij":              true,
		"abcde xyz ecdab":          false,
		"a ab abc abd abf abj":     true,
		"iiii oiii ooii oooi oooo": true,
		"oiii ioii iioi iiio":      false,
	}
	for i, o := range io {
		assert.Equal(t, isValid2(i), o)
	}
}
