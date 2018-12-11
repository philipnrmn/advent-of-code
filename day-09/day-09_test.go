package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	i := []string{"9", "10", "13", "17", "21", "30"}
	i2 := []string{"25", "1618", "7999", "1104", "6111", "5807"}
	o := []int{32, 8317, 146373, 2764, 54718, 37305}
	for n := range i {
		if n > 0 {
			break
		}
		assert.Equal(t, o[n], solve1(i[n], i2[n]))
	}
}
