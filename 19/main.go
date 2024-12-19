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

func parseInput(input string) (patterns []string, desires []string) {
	allLines := strings.Split(input, "\n")
	patterns = strings.Split(allLines[0], ", ")

	for _, d := range allLines[1:] {
		if len(d) > 0 {
			desires = append(desires, d)
		}
	}
	return patterns, desires
}

func findStripe(desired string, patterns []string) (bool, int) {
	allResult := [][]string{}
	all_patterns := map[string]bool{}
	for _, p := range patterns {
		all_patterns[p] = true
	}
	queue := []Towel{{
		desired,
		all_patterns,
		[]string{},
	}}

	for len(queue) > 0 {
		this_towel := queue[0]
		queue = queue[1:]

		if len(this_towel.desired) == 0 {
			allResult = append(allResult, this_towel.result_patterns)
			continue
		}

		for p, ok := range this_towel.patterns {
			if len(p) > len(this_towel.desired) || !ok {
				continue
			}
			if !strings.Contains(this_towel.desired, p) {
				this_towel.patterns[p] = false
			}
			if p == this_towel.desired[:len(p)] {
				queue = append(queue, Towel{
					this_towel.desired[len(p):],
					this_towel.patterns,
					append(this_towel.result_patterns, p),
				})
			}
		}
	}

	if len(allResult) != 0 {
		fmt.Println(allResult)
		return true, len(allResult)
	}
	return false, 0
}

func main() {
	allInput := input
	// allInput := realInput
	patterns, desires := parseInput(allInput)
	// fmt.Println(patterns, desires)

	cnt := 0
	cnt2 := 0
	for _, d := range desires {
		if ok, r := findStripe(d, patterns); ok {
			fmt.Println("found", d, r)
			cnt++
			cnt2 += r
		}
	}
	fmt.Println(cnt)
	fmt.Println(cnt2)
}
