package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`
	assert.Equal(t, solve1(input), "tknk")
}

func TestWeightOf(t *testing.T) {
	input := `pbga (66)
	xhth (57)
	ebii (61)
	havc (66)
	ktlj (57)
	fwft (72) -> ktlj, cntj, xhth
	qoyq (66)
	padx (45) -> pbga, havc, qoyq
	tknk (41) -> ugml, padx, fwft
	jptl (61)
	ugml (68) -> gyxo, ebii, jptl
	gyxo (61)
	cntj (57)`
	nodes := getNodes(input)
	nodeMap := makeNodeMap(nodes)
	io := map[int]int{
		11: 61,
		0:  66,
		4:  57,
		10: 251,
		7:  243,
		5:  243,
	}

	for i, o := range io {
		assert.Equal(t, o, weightOf(nodes[i], nodeMap))
	}
}
