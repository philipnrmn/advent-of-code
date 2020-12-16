package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var NYI = errors.New("Not yet implemented")
var NOPE = errors.New("Did not find a solution")

type Solution func([]string) (int, error)
type ISolution func([]int) (int, error)

func Solve(day int, s ...Solution) {
	if len(os.Args) < 2 {
		barf("Usage: go run day-<day>/main.go <part>")
	}

	var part int
	var err error
	if part, err = strconv.Atoi(os.Args[1]); err != nil {
		barf("Bad value for part", err)
	}

	var input []string
	if input, err = readInput(day); err != nil {
		barf(err)
	}

	fmt.Printf("Solving day %d part %d...\n", day, part)

	out, err := s[part-1](input)
	if err != nil {
		barf(err)
	}

	fmt.Println(out)
}

func barf(err ...interface{}) {
	fmt.Fprintln(os.Stderr, err...)
	os.Exit(1)
}

func readInput(day int) ([]string, error) {
	var input []string
	var err error

	var filename string
	cwd, _ := os.Getwd()
	if strings.HasPrefix(filepath.Base(cwd), "day-") {
		filename = "input"
	} else {
		filename = fmt.Sprintf("day-%02d/input", day)
	}

	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return input, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input, nil
}

func Ints(raw []string) []int {
	var input []int

	for _, r := range raw {
		input = append(input, ToInt(r))
	}
	return input
}

func Chunk(raw []string) [][]string {
	var chunks [][]string
	var chunk []string
	for _, s := range raw {
		if len(s) > 0 {
			chunk = append(chunk, s)
		} else if len(chunk) > 0 {
			chunks = append(chunks, chunk)
			chunk = []string{}
		}
	}
	if len(chunk) > 0 {
		return append(chunks, chunk)
	}
	return chunks
}

func CsvInts(raw []string) [][]int {
	var rows [][]int
	for _, r := range raw {
		var row []int
		for _, v := range strings.Split(r, ",") {
			row = append(row, ToInt(v))
		}
		rows = append(rows, row)
	}
	return rows
}

func ToInt(raw string) int {
	result, err := strconv.Atoi(raw)
	if err != nil {
		barf("Could not cast to int", err)
	}
	return result
}
