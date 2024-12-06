package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Pair struct {
	x, y int
}

const (
	U int = iota
	R
	D
	L
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
			}
		}
		for i, c := range line {
			if c == '#' {
				obstacles = append(obstacles,
					Pair{
						len(result) - 1,
						i,
					})
			}
		}
	}
	return result, guard, obstacles
}

func fillx(input []string, orientation int, index int, from int, to int) {
	// fmt.Println("fillx", orientation, index, from, to)
	switch orientation {
	case U, D:
		for i := from; i <= to; i++ {
			input[i] = input[i][:index] + "X" + input[i][index+1:]
		}
	case L, R:
		for i := from; i <= to; i++ {
			input[index] = input[index][:i] + "X" + input[index][i+1:]
		}
	default:
		fmt.Println("Invalid orientation")
	}
	// for _, line := range input {
	// 	fmt.Println(line)
	// }
}

func letsgo(guard Pair, obstacles []Pair, TURN int, input []string) (bool, Pair) {
	candidates := []Pair{}
	switch TURN {
	case U:
		for _, o := range obstacles {
			if o.y == guard.y && o.x < guard.x {
				candidates = append(candidates, o)
			}
		}
		if len(candidates) == 0 {
			fillx(input, U, guard.y, 0, guard.x)
			return false, Pair{0, 0}
		}
		new_guard := Pair{candidates[len(candidates)-1].x + 1, guard.y}
		fillx(input, U, guard.y, new_guard.x, guard.x)
		return true, new_guard
	case D:
		for _, o := range obstacles {
			if o.y == guard.y && o.x > guard.x {
				candidates = append(candidates, o)
			}
		}
		if len(candidates) == 0 {
			fillx(input, D, guard.y, guard.x, len(input)-1)
			return false, Pair{0, 0}
		}
		new_guard := Pair{candidates[0].x - 1, guard.y}
		fillx(input, D, guard.y, guard.x, new_guard.x)
		return true, new_guard
	case L:
		for _, o := range obstacles {
			if o.x == guard.x && o.y < guard.y {
				candidates = append(candidates, o)
			}
		}
		if len(candidates) == 0 {
			fillx(input, L, guard.x, 0, guard.y)
			return false, Pair{0, 0}
		}
		new_guard := Pair{guard.x, candidates[len(candidates)-1].y + 1}
		fillx(input, L, guard.x, new_guard.y, guard.y)
		return true, new_guard
	case R:
		for _, o := range obstacles {
			if o.x == guard.x && o.y > guard.y {
				candidates = append(candidates, o)
			}
		}
		if len(candidates) == 0 {
			fillx(input, R, guard.x, guard.y, len(input[0])-1)
			return false, Pair{0, 0}
		}
		new_guard := Pair{guard.x, candidates[0].y - 1}
		fillx(input, R, guard.x, guard.y, new_guard.y)
		return true, new_guard
	default:
		return false, Pair{0, 0}
	}
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input, guard, obstacles := read_input(file_name)
	// fmt.Println(input)
	fmt.Println(guard)
	// fmt.Println(obstacles)

	// Sort obstacles
	slices.SortFunc(obstacles, func(a Pair, b Pair) int {
		if a.x == b.x {
			return cmp.Compare(a.y, b.y)
		}
		return cmp.Compare(a.x, b.x)
	})
	var UD []Pair
	for _, o := range obstacles {
		UD = append(UD, o)
	}
	slices.SortFunc(obstacles, func(a Pair, b Pair) int {
		if a.y == b.y {
			return cmp.Compare(a.x, b.x)
		}
		return cmp.Compare(a.y, b.y)
	})
	var LR []Pair
	for _, o := range obstacles {
		LR = append(LR, o)
	}
	// fmt.Println(UD)
	// fmt.Println(LR)

	TURN := U
	for {
		var canGo bool
		switch TURN {
		case U, D:
			canGo, guard = letsgo(guard, UD, TURN, input)
		case L, R:
			canGo, guard = letsgo(guard, LR, TURN, input)
		default: // Should not happen
			log.Fatal("Invalid TURN")
		}
		fmt.Println(canGo, guard)
		if !canGo {
			break
		}
		TURN = (TURN + 1) % 4
	}

	after_cnt := 0
	for _, line := range input {
		after_cnt += strings.Count(line, "X")
		// fmt.Println(line)
	}
	fmt.Println(after_cnt)
}
