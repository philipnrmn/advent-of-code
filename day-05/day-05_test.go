package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	i := `dabAcCaCBAcCcaDA`
	o := 10
	assert.Equal(t, o, solve1(i))
}
