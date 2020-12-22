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

type StringSet struct {
	contents map[string]bool
}

func NewStringSet(ss ...string) *StringSet {
	contents := map[string]bool{}
	for _, s := range ss {
		contents[s] = true
	}
	return &StringSet{contents}
}

func (set *StringSet) Append(ss ...string) {
	for _, s := range ss {
		set.contents[s] = true
	}
}

func (set *StringSet) Delete(ss ...string) {
	for _, s := range ss {
		delete(set.contents, s)
	}
}

func (set *StringSet) Contains(s string) bool {
	return set.contents[s]
}

func (set *StringSet) Len() int {
	return len(set.contents)
}

func (set *StringSet) ToSlice() []string {
	result := make([]string, len(set.contents))
	i := 0
	for k, _ := range set.contents {
		result[i] = k
		i++
	}
	return result
}

func IntersectStringSet(a, b *StringSet) *StringSet {
	contents := map[string]bool{}
	for k, v := range a.contents {
		if v && b.contents[k] {
			contents[k] = true
		}
	}
	return &StringSet{contents}
}

func UnionStringSet(a, b *StringSet) *StringSet {
	contents := map[string]bool{}
	for k, v := range a.contents {
		contents[k] = v
	}
	for k, v := range b.contents {
		contents[k] = v
	}
	return &StringSet{contents}
}
