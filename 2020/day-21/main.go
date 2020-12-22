package main

import (
	"fmt"
	"strings"

	. "../lib"
)

func part1(input []string) (int, error) {
	ingredients, allergens := parse(input)
	almap := mapAllergens(ingredients, allergens)

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
	ingredients, allergens := parse(input)
	almap := mapAllergens(ingredients, allergens)

	target := len(almap)
	result := map[string]string{}
	for {
		stable := true
		for a, ing := range almap {
			if ing.Len() == 1 {
				result[a] = ing.ToSlice()[0]
				for i := range almap {
					almap[i].Delete(result[a])
				}
				stable = false
			}
		}
		if len(result) == target {
			break
		}
		if stable {
			return 0, NOPE
		}
	}

	// run this through sort and sed
	for a, ing := range result {
		fmt.Println(a, ing)
	}

	return 0, nil
}

func mapAllergens(ingredients, allergens [][]string) map[string]*StringSet {
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
	return almap
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
