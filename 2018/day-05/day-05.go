package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
}

func solve1(input string) int {
	runes := []rune(input)
	for i := 1; i < len(runes); i++ {
		switch {
		case unicode.IsLower(runes[i]) && runes[i-1] == unicode.ToUpper(runes[i]):
			fallthrough
		case unicode.IsUpper(runes[i]) && runes[i-1] == unicode.ToLower(runes[i]):
			return solve1(string(append(runes[0:i-1], runes[i+1:]...)))
		}
	}
	return len(input)
}

func solve2(input string) int {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	shortest := len(input)
	var winning rune
	for _, c := range alphabet {
		inp := without(input, c)
		res := solve1(inp)
		if res < shortest {
			shortest = res
			winning = c
		}
	}
	fmt.Printf("Winner is %#U with %d chars", winning, shortest)
	return shortest
}

func without(input string, without rune) string {
	result := []rune{}
	lower := unicode.ToLower(without)
	upper := unicode.ToUpper(without)
	for _, c := range input {
		if c != lower && c != upper {
			result = append(result, c)
		}
	}
	return string(result)
}
