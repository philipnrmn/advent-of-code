package main

import (
	"strings"

	. "../lib"
)

func part1(input []string) (int, error) {
	ingredients, allergens := parse(input)
	almap := map[string]*StringSet{}
	for i := range allergens {
		for _, ag := range allergens[i] {
			if _, ok := almap[ag]; !ok {
				almap[ag] = NewStringSet(ingredients[i]...)
			} else {
				almap[ag] = IntersectStringSet(almap[ag], NewStringSet(ingredients[i]...))
			}
		}
	}

	alset := NewStringSet()
	for _, ing := range almap {
		alset = UnionStringSet(alset, ing)
	}

	count := 0
	for _, i := range ingredients {
		for _, j := range i {
			if !alset.Contains(j) {
				count++
			}
		}
	}

	return count, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func parse(input []string) ([][]string, [][]string) {
	var ingredients [][]string
	var allergens [][]string
	for _, line := range input {
		p := strings.Trim(line, ")")
		q := strings.Split(p, " (contains ")
		ingredients = append(ingredients, strings.Split(q[0], " "))
		allergens = append(allergens, strings.Split(q[1], ", "))
	}
	return ingredients, allergens
}

func main() {
	Solve(21, part1, part2)
}
