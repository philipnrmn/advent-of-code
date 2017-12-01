package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(captcha string) int {
	l := len(captcha)

	first := getIntAt(captcha, 0)
	latest := first
	sum := 0

	for i := 1; i < l; i++ {
		c := getIntAt(captcha, i)
		if latest == c {
			sum += c
		}
		latest = c
	}

	if latest == first {
		sum += latest
	}

	return sum
}

// getIntAt is a helper which returns the integer value of the character
// at index i in string s
func getIntAt(s string, i int) int {
	n, _ := strconv.Atoi(string(s[i]))
	return n
}
