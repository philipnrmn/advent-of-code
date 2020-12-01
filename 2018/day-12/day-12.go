package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(solve1(os.Args[1], os.Args[2], 20))
	fmt.Println(solve2(os.Args[1], os.Args[2], 50000000000))
}

func solve1(i1, i2 string, gen int) int {
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

	for i := 0; i < gen; i++ {
		fmt.Printf("%2d: ", i)
		printGen(state, -2, 32)
		zero += 2
		state = iterate(rules, state)
	}
	fmt.Printf("%2d: ", gen)
	printGen(state, -2, 32)

	return sum(state, zero)
}

func solve2(i1, i2 string, gen int) int {
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

	for i := 0; i < gen; i++ {
		// fmt.Printf("%2d: ", i)
		// printGen(state, -zero, 32)
		fmt.Println(i, " score: ", sum(state, zero))
		zero += 2
		state = iterate(rules, state)
		state, zero = trim(state, zero)
	}
	fmt.Printf("%2d: ", gen)
	printGen(state, -2, 32)

	return sum(state, zero)
}

func trim(state []bool, zero int) ([]bool, int) {
	var start, end int
	for i := range state {
		if state[i] {
			start = i
			break
		}
	}
	for i := len(state) - 1; i >= 0; i-- {
		if state[i] {
			end = i + 1
			break
		}
	}
	return state[start:end], zero - start
}

func iterate(rules map[byte]bool, state []bool) []bool {
	result := make([]bool, len(state)+4)
	for i := -2; i < len(state)+2; i++ {
		b := byteAt(state, i)
		if live, ok := rules[b]; ok && live {
			result[i+2] = true
		}
	}
	return result
}

func sum(state []bool, zero int) int {
	var result int
	for i, b := range state {
		if b {
			result += i - zero
		}
	}
	return result
}

func strtob(s string) byte {
	var result byte
	for _, r := range s {
		result = result << 1
		if r == '#' {
			result = result | 1
		}
	}
	return result
}

func byteAt(state []bool, index int) byte {
	var result byte
	for i := index - 2; i <= index+2; i++ {
		result = result << 1
		if i >= 0 && i < len(state) && state[i] {
			result = result | 1
		}
	}
	return result
}

func printGen(state []bool, start, stop int) {
	var buf string
	for i := start; i < stop; i++ {
		if i >= 0 && i < len(state) && state[i] {
			buf += "#"
		} else {
			buf += "."
		}
	}
	fmt.Println(buf)
}

func prepend(l []bool, n bool) []bool {
	return append([]bool{n}, l...)
}
