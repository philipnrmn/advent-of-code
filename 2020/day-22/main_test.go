package main

import (
	"testing"
	"time"
)

var i = []string{
	"Player 1:",
	"9",
	"2",
	"6",
	"3",
	"1",
	"",
	"Player 2:",
	"5",
	"8",
	"4",
	"7",
	"10",
}

func Test_part1(t *testing.T) {
	e := 306
	o, err := part1(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Test_part2(t *testing.T) {
	e := 291
	o, err := part2(i)
	if err != nil {
		t.Error(err)
	}
	if o != e {
		t.Errorf("Expected %d, got %d", e, o)
	}
}

func Test_recurse(t *testing.T) {
	timeout := time.After(100 * time.Millisecond)
	done := make(chan bool)
	go func() {
		recurse([]int{43, 19}, []int{2, 29, 14})
		done <- true
	}()
	select {
	case <-timeout:
		t.Error("Timed out")
	case <-done:
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
