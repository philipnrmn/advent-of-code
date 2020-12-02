package main

import (
	"regexp"
	"strconv"

	. "../lib"
)

var re = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func part1(input []string) (int, error) {
	valid := 0

	for _, line := range input {
		min, max, char, pwd := parse(line)
		c := count(pwd)
		if c[char] >= min && c[char] <= max {
			valid++
		}
	}
	return valid, nil
}

func part2(input []string) (int, error) {
	valid := 0

	for _, line := range input {
		a, b, char, pwd := parse(line)
		rpwd := []rune(pwd)

		if xor(rpwd[a-1] == char, rpwd[b-1] == char) {
			valid++
		}
	}
	return valid, nil
}

func parse(input string) (int, int, rune, string) {
	groups := re.FindStringSubmatch(input)
	a, _ := strconv.Atoi(groups[1])
	b, _ := strconv.Atoi(groups[2])
	c := []rune(groups[3])[0]
	d := groups[4]

	return a, b, c, d
}

func count(input string) map[rune]int {
	result := map[rune]int{}

	for _, r := range input {
		if _, ok := result[r]; ok {
			result[r]++
		} else {
			result[r] = 1
		}
	}
	return result
}

func xor(a, b bool) bool {
	return (a && !b) || (b && !a)
}

func main() {
	Solve(2, part1, part2)
}
