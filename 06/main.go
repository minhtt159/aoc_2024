package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Pair struct {
	x int
	y int
	d int
}

var (
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, 1, 0, -1}
)

func read_input(file_name string) ([]string, Pair, []Pair) {
	result := []string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var guard Pair
	var obstacles []Pair

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)

		if strings.Contains(line, "^") {
			guard = Pair{
				len(result) - 1,
				strings.Index(line, "^"),
				0,
			}
		}
		for i, c := range line {
			if c == '#' {
				obstacles = append(obstacles,
					Pair{
						len(result) - 1,
						i,
						-1,
					})
			}
		}
	}
	return result, guard, obstacles
}

func fillx(input []string, row int, col int, c rune) {
	input[row] = input[row][:col] + string(c) + input[row][col+1:]
}

func main() {
	file_name := "input.txt"
	// file_name := "real_input.txt"
	input, guard, obstacles := read_input(file_name)
	// fmt.Println(input)
	// fmt.Println(guard)
	// fmt.Println(obstacles)

	TURN := 0
	possible_obstacles := 0
	for {
		// Try do move the guard
		new_x := guard.x + dir_x[TURN]
		new_y := guard.y + dir_y[TURN]
		if new_x < 0 || new_x >= len(input) || new_y < 0 || new_y >= len(input[0]) {
			// Guard moved out of the map, exit
			break
		}
		if slices.Contains(obstacles, Pair{new_x, new_y, -1}) {
			// Guard meet an obstacles, pivot
			TURN = (TURN + 1) % 4
			continue
		}

		// Guard can move, commit the move
		guard = Pair{new_x, new_y, TURN}
		fillx(input, guard.x, guard.y, 'X')

		// Try to place an obstacle in the front
		obstacles = append(obstacles,
			Pair{
				guard.x + dir_x[TURN],
				guard.y + dir_y[TURN],
				-1,
			})

		// Pivot to see if there is any loop can be achieved
		trial_turn := (TURN + 1) % 4
		dummy_guard := Pair{
			guard.x,
			guard.y,
			guard.d,
		}
		dummy_stack := []Pair{dummy_guard}
		can_place_obs := false

		for {
			forward_x := dummy_guard.x + dir_x[trial_turn]
			forward_y := dummy_guard.y + dir_y[trial_turn]

			if forward_x < 0 || forward_x >= len(input) || forward_y < 0 || forward_y >= len(input[0]) {
				// Look ahead but out of map, break
				break
			}

			if slices.Contains(obstacles, Pair{forward_x, forward_y, -1}) {
				// Look ahead but hit an obstacle, pivot and continue
				trial_turn = (trial_turn + 1) % 4
				continue
			}

			// Try the move
			dummy_guard = Pair{
				forward_x,
				forward_y,
				trial_turn,
			}
			// Check if this step is already committed
			if slices.Contains(dummy_stack, dummy_guard) {
				can_place_obs = true
				fmt.Println(dummy_stack)
				fmt.Println("Loop detected, place obstacles at", guard.x+dir_x[TURN], guard.y+dir_y[TURN])
				for _, line := range input {
					fmt.Println(line)
				}
				break
			}
			// Commit the move
			dummy_stack = append(dummy_stack, dummy_guard)
		}
		if can_place_obs {
			possible_obstacles++
		}

		// Remove the obstacle
		obstacles = obstacles[:len(obstacles)-1]
	}

	after_cnt := 0
	for _, line := range input {
		after_cnt += len(line) - strings.Count(line, "#") - strings.Count(line, ".")
		// fmt.Println(line)
	}
	fmt.Println(after_cnt)

	fmt.Println(possible_obstacles)

	// daySix()
}
