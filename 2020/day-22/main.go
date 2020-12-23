package main

import (
	"reflect"
	"strconv"

	. "../lib"
)

func part1(input []string) (int, error) {
	p1, p2 := parse(input)
	for {
		if len(p1) == 0 || len(p2) == 0 {
			break
		}
		p1, p2 = play(p1, p2)
	}

	if len(p2) > len(p1) {
		return score(p2), nil
	}
	return score(p1), nil
}

func part2(input []string) (int, error) {
	p1, p2 := parse(input)
	p1, p2 = recurse(p1, p2)

	if len(p2) > len(p1) {
		return score(p2), nil
	}
	return score(p1), nil
}

func recurse(p1, p2 []int) ([]int, []int) {
	cache := [][][]int{}
	for {
		round := [][]int{p1, p2}
		for i := range cache {
			if reflect.DeepEqual(round, cache[i]) {
				return p1, []int{}
			}
		}
		cache = append(cache, round)

		if len(p1) == 0 || len(p2) == 0 {
			break
		}
		if len(p1) > p1[0] && len(p2) > p2[0] {
			// I learned a painful lesson about slice internals here
			p1a := make([]int, p1[0])
			p2a := make([]int, p2[0])
			copy(p1a, p1[1:p1[0]+1])
			copy(p2a, p2[1:p2[0]+1])

			p1a, p2a = recurse(p1a, p2a)
			if len(p1a) > len(p2a) {
				p1 = append(p1[1:], p1[0], p2[0])
				p2 = p2[1:]
			} else {
				p2 = append(p2[1:], p2[0], p1[0])
				p1 = p1[1:]
			}
		} else {
			p1, p2 = play(p1, p2)
		}
	}
	return p1, p2
}

func play(p1, p2 []int) ([]int, []int) {
	// hopefully we never draw?
	if p1[0] > p2[0] {
		p1 = append(p1[1:], p1[0], p2[0])
		p2 = p2[1:]
	} else {
		p2 = append(p2[1:], p2[0], p1[0])
		p1 = p1[1:]
	}
	return p1, p2
}

func score(deck []int) int {
	result := 0
	for i := 1; i <= len(deck); i++ {
		result += deck[len(deck)-i] * i
	}
	return result
}

func parse(input []string) ([]int, []int) {
	chunks := Chunk(input)
	decks := make([][]int, 2)
	for i := 0; i < 2; i++ {
		decks[i] = make([]int, len(chunks[0])-1)
		for j := 0; j < len(decks[i]); j++ {
			decks[i][j], _ = strconv.Atoi(chunks[i][j+1])
		}
	}
	return decks[0], decks[1]
}

func main() {
	Solve(22, part1, part2)
}
