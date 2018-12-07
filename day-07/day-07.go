package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(solve1(os.Args[1]))
}

func solve1(input string) string {
	deps := map[string][]string{}
	ids := []string{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		id := words[7]
		dep := words[1]
		ids = addToSet(ids, id)
		ids = addToSet(ids, dep)
		deps[id] = append(deps[id], dep)
	}

	sort.Strings(ids)

	sorted := make([]string, 0, len(ids))
	for {
		for i := 0; i < len(ids); i++ {
			id := ids[i]
			if !contains(sorted, id) && containsSlice(sorted, deps[id]) {
				sorted = append(sorted, id)
				break
			}
		}
		if len(sorted) == len(ids) {
			break
		}
	}
	return strings.Join(sorted, "")
}

func addToSet(set []string, item string) []string {
	if !contains(set, item) {
		return append(set, item)
	}
	return set
}

func contains(haystack []string, needle string) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}

// I only called them that to make Fabs happy
func containsSlice(chicken, egg []string) bool {
	for _, e := range egg {
		if !contains(chicken, e) {
			return false
		}
	}
	return true
}
