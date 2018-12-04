package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) int {
	records := strings.Split(input, "\n")

	const elftime = "[2006-01-02 15:04]"

	change := regexp.MustCompile("^(.+) Guard #(\\d+) begins shift")
	asleep := regexp.MustCompile("^(.+) falls asleep")
	awake := regexp.MustCompile("^(.+) wakes up")

	// sort by timestamp
	sort.Strings(records)
	totals := map[int]int{}
	minutes := map[int]map[int]int{}
	for i := 0; i < 60; i++ {
		minutes[i] = map[int]int{}
	}

	var start time.Time
	var gID int
	for _, r := range records {
		switch {
		case change.MatchString(r):
			matches := change.FindStringSubmatch(r)
			gID, _ = strconv.Atoi(matches[2])
			break
		case asleep.MatchString(r):
			matches := asleep.FindStringSubmatch(r)
			start, _ = time.Parse(elftime, matches[1])
			break
		case awake.MatchString(r):
			matches := awake.FindStringSubmatch(r)
			stop, _ := time.Parse(elftime, matches[1])

			delta := int(stop.Sub(start).Minutes())
			totals[gID] += delta
			for i := 0; i < delta; i++ {
				minutes[(int(start.Minute())+i)%60][gID]++
			}
			break
		default:
			fmt.Printf("Could not parse %s\n", r)
			break
		}
	}

	var best int
	for g, t := range totals {
		if t > best {
			gID = g
			best = t
		}
	}
	fmt.Printf("Guard #%d slept longest, at %d minutes\n", gID, best)

	best = 0
	var mID int
	for i := 0; i < 60; i++ {
		m := minutes[i][gID]
		if m > best {
			mID = i
			best = m
		}
	}
	fmt.Printf("Minute %d was the most common, at %d times\n", mID, best)
	return gID * mID
}
