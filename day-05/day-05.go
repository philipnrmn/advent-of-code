package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
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
