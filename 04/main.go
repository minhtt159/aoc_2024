package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func read_input(file_name string) []string {
	result := []string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

type Pair struct {
	a, b interface{}
}

var (
	dir_x    = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dir_y    = []int{-1, 0, 1, -1, 1, -1, 0, 1}
	target_0 = "XMAS"
	target_1 = []string{"MAS", "SAM"}
)

// Get a map of words from the input
// BFS search all 8 directions for the word "XMAS"
// Return the number of times the word "XMAS" appears
func search(puzzle []string, x int, y int, directions Pair, buf string, target string) int {
	if x < 0 || y < 0 || x >= len(puzzle) || y >= len(puzzle[0]) {
		return 0
	}
	current := buf + string(puzzle[x][y])
	if current == target {
		return 1
	}
	// if len(current) == len(target) {
	// 	return 0
	// }
	if strings.Contains(target, current) {
		return search(puzzle,
			x+directions.a.(int),
			y+directions.b.(int),
			directions,
			current,
			target,
		)
	}
	return 0
}

func main01(input []string) {
	total := 0
	for x, line := range input {
		for y, c := range line {
			if c != 'X' {
				continue
			}
			result := 0
			for d := 0; d < 8; d++ {
				this_round := search(input,
					x+dir_x[d],
					y+dir_y[d],
					Pair{dir_x[d], dir_y[d]},
					"X",
					target_0,
				)
				// if this_round != 0 {
				// 	fmt.Println(x, y, dir_x[d], dir_y[d])
				// }
				result += this_round
			}
			total += result
		}
	}
	fmt.Println(total)
}

func main02(input []string) {
	total := 0
	for x, line := range input {
		for y, c := range line {
			if c != 'A' {
				continue
			}
			if x < 1 || y < 1 || x > len(input)-2 || y > len(input[0])-2 {
				continue
			}
			edge_1 := string(input[x-1][y-1]) + "A" + string(input[x+1][y+1])
			edge_2 := string(input[x-1][y+1]) + "A" + string(input[x+1][y-1])
			if slices.Contains(target_1, edge_1) && slices.Contains(target_1, edge_2) {
				total++
			}
		}
	}
	fmt.Println(total)
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	main01(input)
	main02(input)
}
