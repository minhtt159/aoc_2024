package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stage struct {
	combo string
	cur   int
}

var round_cache = make(map[Stage]int)

func getMinPath(combo string, cur int, max int) int {
	// combo: combination of keys in the dir pad
	// cur: current round
	// max: maximum round
	if val, ok := round_cache[Stage{combo, cur}]; ok {
		return val
	}
	result := 0
	start := Pos{0, 2} // A of dir map
	for _, c := range combo {
		end := findPosInMap(dir_map, c)
		if cur == max {
			// Just pick a random one since all the length are the same
			result += len(cache[Cache{0, start, end}][0])
		} else {
			possible_path := cache[Cache{0, start, end}]
			shortest_path := math.MaxInt64
			// Get the shortest path
			for _, path := range possible_path {
				temp := getMinPath(path, cur+1, max)
				if shortest_path > temp {
					shortest_path = temp
				}
			}
			result += shortest_path
		}
		start = end
	}
	// fmt.Println("Cache", combo, cur, "=>", result)
	round_cache[Stage{combo, cur}] = result
	return result
}

func main() {
	// Init cache
	for i, line1 := range num_map {
		for j, r1 := range line1 {
			for k, line2 := range num_map {
				for v, r2 := range line2 {
					if r1 == N || r2 == N {
						// useless
						continue
					}
					Dijkstra(num_map, Pos{i, j}, Pos{k, v})
				}
			}
		}
	}
	for i, line1 := range dir_map {
		for j, r1 := range line1 {
			for k, line2 := range dir_map {
				for v, r2 := range line2 {
					if r1 == N || r2 == N {
						// useless
						continue
					}
					Dijkstra(dir_map, Pos{i, j}, Pos{k, v})
				}
			}
		}
	}
	// Solve
	xxx := case4
	input, round := xxx.input, xxx.round
	combos := strings.Split(strings.Trim(input, "\n"), "\n")
	sum := 0

	// First numpad
	for _, combo := range combos {
		weight, _ := strconv.Atoi(strings.Trim(combo, "A"))
		stage_1 := []string{""}
		start := Pos{3, 2} // A of num map
		for _, r := range combo {
			end := findPosInMap(num_map, r)

			possible_path := cache[Cache{1, start, end}]
			next_step := []string{}

			for _, path := range possible_path {
				for _, prev_path := range stage_1 {
					next_step = append(next_step, prev_path+path)
				}
			}

			start = end
			stage_1 = next_step
		}
		fmt.Println("Path", combo, weight, "=>", stage_1)
		// Direction pad
		length := math.MaxInt64
		for _, line := range stage_1 {
			temp := getMinPath(line, 1, round)
			if temp < length {
				length = temp
			}
		}

		sum += length * weight
		// break
	}

	fmt.Println("Sum", sum)
}
