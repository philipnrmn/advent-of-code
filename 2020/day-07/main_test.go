package main

import (
	"reflect"
	"testing"
)

var i = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

func Test_parse(t *testing.T) {
	eOuter := "light red"
	eInner := map[string]int{
		"bright white": 1,
		"muted yellow": 2,
	}
	outer, inner := parse(i[0])
	if outer != eOuter {
		t.Errorf("Expected %s, got %s", eOuter, outer)
	}
	if !reflect.DeepEqual(inner, eInner) {
		t.Errorf("Expected %v, got %v", eInner, inner)
	}
}

func Test_part1(t *testing.T) {
	e := 4
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Test_part2(t *testing.T) {
	e := 32
	o, err := part2(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Benchmark_part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part1(i)
	}
}

func Benchmark_part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2(i)
	}
}
