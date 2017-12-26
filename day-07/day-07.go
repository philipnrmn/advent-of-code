package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// don't actually make trees like this
type node struct {
	name     string
	weight   int
	children []string
}

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) string {
	nodes := getNodes(input)
	children := map[string]bool{}

	for _, n := range nodes {
		for _, c := range n.children {
			children[c] = true
		}
	}

	for _, n := range nodes {
		if _, found := children[n.name]; !found {
			return n.name
		}
	}
	return "ERR"
}

func getNodes(input string) []node {
	lines := strings.Split(input, "\n")
	nodes := make([]node, len(lines))
	for i, l := range lines {
		nodes[i] = stringToNode(l)
	}
	return nodes
}

// stringToNode extracts a node from its serialized format
func stringToNode(line string) node {
	re := regexp.MustCompile(`(\w+) \((\d+)\)(?: -> )?([\w, ]*)`)
	sm := re.FindAllStringSubmatch(line, -1)
	// This warning just helps me remember the dquotes in
	// go run day-07.go "$(</tmp/input)"
	if sm == nil || len(sm) == 0 {
		fmt.Println("Badly formatted:", line)
		return node{}
	}
	w, _ := strconv.Atoi(sm[0][2])
	return node{
		name:     sm[0][1],
		weight:   w,
		children: strings.Split(sm[0][3], ", "),
	}
}
