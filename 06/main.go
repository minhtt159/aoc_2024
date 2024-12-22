package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

type Pos struct {
	x int
	y int
	d int
}

var (
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, 1, 0, -1}
)

func read_input(input string) ([][]byte, Pos, []Pos) {
	var guard Pos
	obstacles := []Pos{}
	all_lines := strings.Split(input, "\n")
	result := make([][]byte, len(all_lines)-1)

	for i, line := range all_lines {
		if len(line) == 0 {
			continue
		}
		result[i] = make([]byte, len(line))
		for j, r := range line {
			result[i][j] = byte(r)
			switch r {
			case '#':
				obstacles = append(obstacles,
					Pos{
						i,
						j,
						-1,
					})
			case '^':
				guard = Pos{
					i,
					j,
					0,
				}
			}
		}
	}
	return result, guard, obstacles
}

func walk(input [][]byte, guard Pos, obstacles []Pos) (bool, []Pos) {
	this_guard := Pos{
		guard.x,
		guard.y,
		guard.d,
	}
	path := []Pos{this_guard}
	for {
		turn := this_guard.d
		new_guard := Pos{
			this_guard.x + dir_x[turn],
			this_guard.y + dir_y[turn],
			turn,
		}

		if new_guard.x < 0 || new_guard.x >= len(input) || new_guard.y < 0 || new_guard.y >= len(input[0]) {
			// fmt.Println("Out of bound, exit")
			break
		}

		if slices.Contains(path, Pos{new_guard.x, new_guard.y, turn}) {
			// fmt.Println("Loop detected, exit")
			return false, []Pos{}
		}

		if slices.Contains(obstacles, Pos{new_guard.x, new_guard.y, -1}) {
			// fmt.Println("Obstacle detected, pivot")
			new_guard = Pos{
				this_guard.x,
				this_guard.y,
				(this_guard.d + 1) % 4,
			}
		} else {
			// Normal move, continue
			input[new_guard.x][new_guard.y] = 'X'
		}

		path = append(path, new_guard)
		this_guard = new_guard
	}
	// fmt.Println("path", path)

	result := []Pos{}
	for _, p := range path {
		if slices.Contains(result, Pos{p.x, p.y, -1}) {
			continue
		}
		result = append(result, Pos{p.x, p.y, -1})
	}

	return true, result
}

func debug(input [][]byte, path []Pos) {
	for i, line := range input {
		for j, r := range line {
			if slices.Contains(path, Pos{i, j, -1}) {
				fmt.Print("X")
			} else {
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}
}

func main() {
	// file_name := ex_input
	file_name := real_input
	input, guard, obstacles := read_input(file_name)
	// fmt.Println(input)
	// fmt.Println(guard)
	// fmt.Println(obstacles)

	_, path := walk(input, guard, obstacles)
	// debug(input, path)
	fmt.Println("p1", len(path))

	p2 := 0
	for i := 1; i < len(path); i++ {
		// if i%100 == 0 {
		// 	fmt.Println(i)
		// }
		dummy_obs := Pos{
			path[i].x,
			path[i].y,
			-1,
		}
		new_obstacles := append(obstacles, dummy_obs)
		// fmt.Println("New obstacle: ", new_obstacle)
		valid, _ := walk(input, guard, new_obstacles)
		if !valid {
			p2++
		}
		// fmt.Println(guard)
	}
	fmt.Println("p2", p2)
}
