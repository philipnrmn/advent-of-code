package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
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

func BenchmarkWithoutTrim(b *testing.B) {
	i1 := `...## => #
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
	i2 := `#..#.#..##......###...###`
	rules := make(map[byte]bool)
	for _, s := range strings.Split(i1, "\n") {
		pattern := strtob(s[0:5])
		rules[pattern] = s[9] == '#'
	}

	state := make([]bool, len(i2))
	for i, r := range i2 {
		state[i] = r == '#'
	}

	// track the zeroth element (state will grow!)
	zero := 0

	for i := 0; i < b.N; i++ {
		zero += 2
		state = iterate(rules, state)
	}
}

func BenchmarkWithTrim(b *testing.B) {
	i1 := `...## => #
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
	i2 := `#..#.#..##......###...###`
	rules := make(map[byte]bool)
	for _, s := range strings.Split(i1, "\n") {
		pattern := strtob(s[0:5])
		rules[pattern] = s[9] == '#'
	}

	state := make([]bool, len(i2))
	for i, r := range i2 {
		state[i] = r == '#'
	}

	// track the zeroth element (state will grow!)
	zero := 0

	for i := 0; i < b.N; i++ {
		zero += 2
		state = iterate(rules, state)
		state, zero = trim(state, zero)
	}
}
