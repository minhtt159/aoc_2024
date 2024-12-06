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

func next_step(input []string, guard Pair, obstacles []Pair, path []Pair) (bool, Pair) {
	turn := guard.d
	new_x := guard.x + dir_x[turn]
	new_y := guard.y + dir_y[turn]

	if new_x < 0 || new_x >= len(input) || new_y < 0 || new_y >= len(input[0]) {
		// fmt.Println("Out of bound, exit")
		return false, Pair{new_x, new_y, -1}
	}

	if slices.Contains(obstacles, Pair{new_x, new_y, -1}) {
		// fmt.Println("Obstacle detected, pivot")
		turn = (turn + 1) % 4
		guard = Pair{guard.x, guard.y, turn}
		return true, guard
	}

	if slices.Contains(path, Pair{new_x, new_y, turn}) {
		// fmt.Println("Loop detected, exit")
		return false, Pair{}
	}

	// Normal move, continue
	guard = Pair{new_x, new_y, turn}
	fillx(input, guard.x, guard.y, 'X')
	return true, guard
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input, guard, obstacles := read_input(file_name)
	// fmt.Println(input)
	// fmt.Println(guard)
	// fmt.Println(obstacles)
	orig_guard := Pair{
		guard.x,
		guard.y,
		guard.d,
	}
	orig_input := []string{}
	for _, line := range input {
		orig_input = append(orig_input, line)
	}
	// fmt.Println(orig_input)

	path := []Pair{}
	var err bool
	for {
		err, guard = next_step(input, guard, obstacles, path)
		if !err {
			break
		}
		// fmt.Println(guard)
		path = append(path, guard)
	}
	// fmt.Println(path)

	after_cnt := 0
	for _, line := range input {
		after_cnt += len(line) - strings.Count(line, "#") - strings.Count(line, ".")
		// fmt.Println(line)
	}
	fmt.Println(after_cnt)

	possible_obstacles := []Pair{}
	for i := 1; i < len(path); i++ {
		// if i%100 == 0 {
		// 	fmt.Println(i)
		// }
		new_obstacle := path[i]
		// fmt.Println("New obstacle: ", new_obstacle)
		obstacles = append(obstacles,
			Pair{
				new_obstacle.x,
				new_obstacle.y,
				-1,
			})
		guard = orig_guard
		new_input := []string{}
		for _, line := range orig_input {
			new_input = append(new_input, line)
			// fmt.Println(line)
		}
		new_path := []Pair{}
		for {
			err, guard = next_step(new_input, guard, obstacles, new_path)
			if !err {
				if (guard == Pair{}) {
					if !slices.Contains(possible_obstacles, Pair{new_obstacle.x, new_obstacle.y, -1}) {
						possible_obstacles = append(possible_obstacles, Pair{new_obstacle.x, new_obstacle.y, -1})
					}
					// fmt.Println("Possible obstacles: ", path[i])
					// for _, line := range new_input {
					// 	fmt.Println(line)
					// }
				}
				break
			}
			// fmt.Println(guard)
			new_path = append(new_path, guard)
		}
		// Remove the obstacle
		obstacles = obstacles[:len(obstacles)-1]
	}
	fmt.Println(len(possible_obstacles))
}
