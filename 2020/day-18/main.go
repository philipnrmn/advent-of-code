package main

import (
	. "../lib"
)

func part1(input []string) (int, error) {
	var result int
	for _, line := range input {
		tokens := tokenize(line)
		groups := lex(tokens)
		result += eval(groups)
	}

	return result, nil
}

func part2(input []string) (int, error) {
	return 0, NYI
}

func eval(groups [][]string) int {
	var sum int
	for _, group := range groups {
		sign := group[0]
		var n int
		if len(group) == 2 {
			n = ToInt(group[1])
		} else {
			n = eval(lex(group[1:]))
		}
		if sign == "+" {
			sum += n
		} else {
			sum *= n
		}
	}
	return sum
}

func lex(tokens []string) [][]string {
	var groups [][]string
	sign := "+"
	depth := 0
	lpi := 0
	for i, t := range tokens {
		switch t {
		case "+", "*":
			if depth == 0 {
				sign = t
			}
		case "(":
			if depth == 0 {
				lpi = i
			}
			depth++
		case ")":
			depth--
			if depth == 0 {
				groups = append(groups, append([]string{sign}, tokens[lpi+1:i]...))
			}
		default:
			if depth == 0 {
				groups = append(groups, []string{sign, t})
			}
		}
	}
	return groups
}

func tokenize(line string) []string {
	var tokens []string

	t := []rune{}
	for _, c := range line {
		switch c {
		case ' ':
			if len(t) > 0 {
				tokens = append(tokens, string(t))
			}
			t = []rune{}
			continue
		case '(':
			tokens = append(tokens, string(c))
		case ')':
			if len(t) > 0 {
				tokens = append(tokens, string(t))
			}
			tokens = append(tokens, string(c))
			t = []rune{}
		default:
			t = append(t, c)
		}
	}

	if len(t) > 0 {
		tokens = append(tokens, string(t))
	}

	return tokens
}

func main() {
	Solve(18, part1, part2)
}
