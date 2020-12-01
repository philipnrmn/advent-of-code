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
	// fmt.Println(solve1(os.Args[1]))
	fmt.Println(solve2(os.Args[1]))
}

func solve2(input string) string {
	nodes := getNodes(input)
	nodeMap := makeNodeMap(nodes)

	for _, n := range nodes {
		if len(n.children) == 0 {
			continue
		}
		cweight := []int{}
		for _, name := range n.children {
			cn := nodeMap[name]
			cweight = append(cweight, weightOf(cn, nodeMap))
		}
		// This is ugly; I ought to apply the delta to the odd-one-out, but
		// humans are good at that and this works so I just print it instead.
		for _, w := range cweight {
			delta := cweight[0] - w
			if delta != 0 {
				fmt.Println("unbalanced node:", n, "weights: ", cweight)
				for _, name := range n.children {
					cn := nodeMap[name]
					fmt.Println(cn.name, cn.weight)
				}
				return n.name
			}
		}

	}

	return "ERR"
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
	// nil strings are "" in Go, calling Split on a nil will yield [""]
	var cc []string
	if len(sm[0][3]) > 0 {
		cc = strings.Split(sm[0][3], ", ")
	}

	return node{
		name:     sm[0][1],
		weight:   w,
		children: cc,
	}
}

// weightOf returns the weight of the node and all its children
func weightOf(n node, m map[string]node) int {
	if len(n.children) == 0 {
		return n.weight
	}
	cweight := 0
	for _, c := range n.children {
		cweight += weightOf(m[c], m)
	}

	return cweight + n.weight
}

// makeNodeMap maps node names to nodes
func makeNodeMap(nn []node) map[string]node {
	m := map[string]node{}
	for _, n := range nn {
		m[n.name] = n
	}
	return m
}
