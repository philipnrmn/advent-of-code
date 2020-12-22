package main

import (
	"strconv"

	. "../lib"
)

func part1(input []string) (int, error) {
	p1, p2 := parse(input)
	for {
		if len(p1) == 0 || len(p2) == 0 {
			break
		}
		// hopefully we never draw?
		if p1[0] > p2[0] {
			p1 = append(p1[1:], p1[0], p2[0])
			p2 = p2[1:]
		} else {
			p2 = append(p2[1:], p2[0], p1[0])
			p1 = p1[1:]
		}
	}

	winner := p1
	if len(p2) > len(p1) {
		winner = p2
	}

	result := 0
	for i := 1; i <= len(winner); i++ {
		result += winner[len(winner)-i] * i
	}
	return result, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
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
