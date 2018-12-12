package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	rules := `...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`
	state := `#..#.#..##......###...###`
	o := 325
	assert.Equal(t, o, solve1(rules, state, 20))
}

func TestSolve2(t *testing.T) {
	rules := `...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`
	state := `#..#.#..##......###...###`
	o := 325
	assert.Equal(t, o, solve2(rules, state, 20))
}

func TestStrtob(t *testing.T) {
	assert.Equal(t, byte(0x4), strtob("..#.."))
	assert.Equal(t, byte(0x15), strtob("#.#.#"))
}
