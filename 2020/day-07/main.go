package main

import (
	"regexp"
	"strconv"

	. "../lib"
)

var ORE = regexp.MustCompile(`(\w+ \w+) bags contain`)
var IRE = regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

func part1(input []string) (int, error) {
	rules := map[string]map[string]int{}
	for _, line := range input {
		o, ii := parse(line)
		rules[o] = ii
	}

	return len(parents("shiny gold", rules)), nil
}

func part2(input []string) (int, error) {
	rules := map[string]map[string]int{}
	for _, line := range input {
		o, ii := parse(line)
		rules[o] = ii
	}
	return countChildren("shiny gold", rules), nil
}

func parse(input string) (string, map[string]int) {
	oGroups := ORE.FindStringSubmatch(input)
	iGroups := IRE.FindAllStringSubmatch(input, -1)
	outer := oGroups[1]
	inner := map[string]int{}
	for _, g := range iGroups {
		inner[g[2]], _ = strconv.Atoi(g[1])
	}
	return outer, inner
}

func parents(colour string, rules map[string]map[string]int) map[string]bool {
	results := map[string]bool{}
	for o, i := range rules {
		if _, ok := i[colour]; ok {
			results[o] = true
			for k, _ := range parents(o, rules) {
				results[k] = true
			}
		}
	}

	return results
}

func countChildren(colour string, rules map[string]map[string]int) int {
	var sum int
	for o, s := range rules[colour] {
		sum += s + s*countChildren(o, rules)
	}
	return sum
}

func main() {
	Solve(07, part1, part2)
}
