package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	. "../lib"
)

var hcl = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var pid = regexp.MustCompile(`^\d{9}$`)
var hgt = regexp.MustCompile(`^(\d{2,3})(cm|in)$`)
var ecls = map[string]bool{
	"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
}

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
	var count int
	for _, passport := range parse(input) {
		if len(passport) < 7 {
			continue
		}
		if _, ok := passport["cid"]; len(passport) == 7 && ok {
			continue
		}
		byr := toInt(passport["byr"])
		if byr < 1920 || byr > 2002 {
			continue
		}
		iyr := toInt(passport["iyr"])
		if iyr < 2010 || iyr > 2020 {
			continue
		}
		eyr := toInt(passport["eyr"])
		if eyr < 2020 || eyr > 2030 {
			continue
		}
		hgt, unit, err := parseHeight(passport["hgt"])
		if err != nil {
			continue
		}
		if unit == "cm" && (hgt < 150 || hgt > 193) {
			continue
		}
		if unit == "in" && (hgt < 59 || hgt > 76) {
			continue
		}
		if !hcl.MatchString(passport["hcl"]) {
			continue
		}
		if !ecls[passport["ecl"]] {
			continue
		}
		if !pid.MatchString(passport["pid"]) {
			continue
		}

		count++
	}
	return count, nil
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

func parseHeight(raw string) (int, string, error) {
	groups := hgt.FindStringSubmatch(raw)
	if len(groups) == 0 {
		return 0, "", errors.New("invalid height")
	}
	return toInt(groups[1]), groups[2], nil
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	Solve(4, part1, part2)
}
