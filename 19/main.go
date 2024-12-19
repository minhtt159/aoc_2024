package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed real_input.txt
var realInput string

type Towel struct {
	desired         string
	patterns        map[string]bool
	result_patterns []string
}

func parseInput(input string) (map[string]bool, []string) {
	allLines := strings.Split(input, "\n")

	patterns := make(map[string]bool)
	for _, p := range strings.Split(allLines[0], ", ") {
		patterns[p] = true
	}

	desires := make([]string, 0)
	for _, d := range allLines[1:] {
		if len(d) > 0 {
			desires = append(desires, d)
		}
	}
	return patterns, desires
}

var cache = make(map[string]int)

func findStripe(desired string, patterns map[string]bool) (bool, int) {
	if cached, ok := cache[desired]; ok {
		return ok, cached
	} else if len(desired) == 0 {
		// done
		return true, 1
	} else if len(patterns) == 0 {
		// no
		return false, 0
	}

	total := 0
	new_patterns := getNewPattern(patterns)

	for p, ok := range new_patterns {
		if !ok || !strings.Contains(desired, p) {
			delete(patterns, p)
			continue
		}
	}

	// fmt.Println("DEBUG", desired, patterns)

	for p := range new_patterns {
		if len(p) > len(desired) {
			continue
		}
		if p == desired[:len(p)] {
			if new_ok, r := findStripe(desired[len(p):], getNewPattern(patterns)); new_ok {
				total += r
			}
		}
	}
	// fmt.Println("Caching", desired, total)
	cache[desired] = total
	return total != 0, total
}

func getNewPattern(patterns map[string]bool) map[string]bool {
	new_patterns := make(map[string]bool)
	for p, ok := range patterns {
		if !ok {
			continue
		}
		new_patterns[p] = true
	}
	return new_patterns
}

func main() {
	// allInput := input
	allInput := realInput
	patterns, desires := parseInput(allInput)
	// fmt.Println(patterns, desires)

	cnt := 0
	cnt2 := 0
	for _, d := range desires {
		new_patterns := getNewPattern(patterns)
		if ok, r := findStripe(d, new_patterns); ok {
			fmt.Println("found", d, r)
			cnt++
			cnt2 += r
		}
	}
	fmt.Println(cnt)
	fmt.Println(cnt2)
}
