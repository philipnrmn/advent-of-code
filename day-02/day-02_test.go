package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	spreadsheet := `5	1	9	5
7	5	3
2	4	6	8`
	assert.Equal(t, solve1(spreadsheet), 18)
}

func TestSolve2(t *testing.T) {
	spreadsheet := `5	9	2	8
9	4	7	3
3	8	6	5`
	assert.Equal(t, solve2(spreadsheet), 9)
}
