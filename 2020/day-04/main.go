package main

import (
	"strings"

	. "../lib"
)

func part1(input []string) (int, error) {
	var count int
	for _, passport := range parse(input) {
		_, ok := passport["cid"]
		if len(passport) == 8 || (len(passport) == 7 && !ok) {
			count++
		}
	}
	return count, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(raw []string) []map[string]string {
	var results []map[string]string

	result := map[string]string{}
	for _, line := range raw {
		if line == "" {
			results = append(results, result)
			result = map[string]string{}
			continue
		}
		for _, pair := range strings.Split(line, " ") {
			kv := strings.Split(pair, ":")
			result[kv[0]] = kv[1]
		}
	}
	results = append(results, result)
	return results
}

func main() {
	Solve(04, part1, part2)

}
