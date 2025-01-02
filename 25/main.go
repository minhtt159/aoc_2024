package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

func parseLayer(buf []string) ([]int, int) {
	// [0,5,3,4,3]
	// #####
	// .####
	// .####
	// .####
	// .#.#.
	// .#...
	// .....
	// [5,0,2,1,3]
	// .....
	// #....
	// #....
	// #...#
	// #.#.#
	// #.###
	// #####

	kind := 0
	pad := []int{-1, -1, -1, -1, -1}
	for _, line := range buf {
		for i, c := range line {
			if c == '#' {
				pad[i]++
			}
		}
	}
	if buf[0][0] == '#' {
		kind = 0
	} else {
		kind = 1
	}
	return pad, kind
}

func parseInput(input string) ([][]int, [][]int) {
	lock := make([][]int, 0)
	key := make([][]int, 0)

	buf := []string{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) != 0 {
			buf = append(buf, line)
		} else {
			// Convert to height
			if len(buf) != 7 {
				log.Fatalf("Invalid input: %s", buf)
			}
			if pad, kind := parseLayer(buf); kind == 0 {
				key = append(key, pad)
			} else {
				lock = append(lock, pad)
			}
			buf = []string{}
		}
	}
	return lock, key
}

func main() {
	// input := ex_input
	input := real_input

	lock, key := parseInput(input)
	fmt.Println(lock, key)

	total := 0
	for _, l := range lock {
		for _, k := range key {
			canFit := true
			for i := range 5 {
				l_i, k_i := l[i], k[i]
				if l_i+k_i > 5 {
					canFit = false
					break
				}
			}
			if canFit {
				total++
			}
		}
	}
	fmt.Println(total)
}
