package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

var test_input = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<` // 2028

var test_input2 = `########
#.@O.O.#
########

>`

//go:embed input.txt
var ex_input string // 10092

//go:embed real_input.txt
var real_input string

var (
	dir_x = []int{-1, 0, 0, 1}
	dir_y = []int{0, -1, 1, 0}
)

const (
	U = iota
	L
	R
	D
)

type Pos struct {
	x int
	y int
}

func parseInput(input string) ([][]byte, Pos, []int) {
	lines := strings.Split(input, "\n")
	warehouse := make([][]byte, len(lines))
	robot := Pos{-1, -1}
	for i, line := range lines {
		if len(line) == 0 {
			warehouse = warehouse[:i]
			break
		}
		warehouse[i] = make([]byte, len(line))
		for j, c := range line {
			warehouse[i][j] = byte(c)
		}
		if strings.Contains(line, "@") {
			robot = Pos{i, strings.Index(line, "@")}
		}
	}
	instructions := []int{}
	for i := 0; i < (len(lines) - len(warehouse)); i++ {
		line := lines[i+len(warehouse)]
		for _, c := range line {
			switch c {
			case '^':
				instructions = append(instructions, U)
			case 'v':
				instructions = append(instructions, D)
			case '<':
				instructions = append(instructions, L)
			case '>':
				instructions = append(instructions, R)
			}
		}
	}
	return warehouse, robot, instructions
}

func getSum(warehouse [][]byte) int {
	sum := 0
	// fmt.Println("DEBUG")
	for i, line := range warehouse {
		for j, c := range line {
			if c == 'O' {
				sum += 100*i + j
			}
		}
		// fmt.Println(string(line))
	}
	// fmt.Println("END")
	return sum
}

func seekMove(warehouse [][]byte, start Pos, direction int) (bool, Pos) {
	if start.x < 0 || start.x >= len(warehouse) || start.y < 0 || start.y >= len(warehouse[0]) {
		// out of map
		return false, Pos{-1, -1}
	}
	next_x := start.x + dir_x[direction]
	next_y := start.y + dir_y[direction]
	switch warehouse[next_x][next_y] {
	case '#':
		// wall
		return false, Pos{-1, -1}
	case '.':
		// empty pos
		return true, Pos{next_x, next_y}
	case 'O':
		return seekMove(warehouse, Pos{next_x, next_y}, direction)
	}
	log.Fatal("Unknown object")
	return false, Pos{-1, -1}
}

func rotateMove(warehouse [][]byte, start Pos, end Pos, direction int) Pos {
	if start.x == end.x {
		// horizontal move
		x := start.x
		origin_y := start.y
		for {
			next_y := end.y - dir_y[direction]
			warehouse[x][end.y] = warehouse[x][next_y]
			if next_y == start.y {
				warehouse[x][origin_y] = '.'
				break
			}
			end.y = next_y
		}
		return Pos{x, origin_y + dir_y[direction]}
	} else {
		// vertical move
		y := start.y
		origin_x := start.x
		for {
			next_x := end.x - dir_x[direction]
			warehouse[end.x][y] = warehouse[next_x][y]
			if next_x == start.x {
				warehouse[origin_x][y] = '.'
				break
			}
			end.x = next_x
		}
		return Pos{origin_x + dir_x[direction], y}
	}
}

func checkMove(warehouse [][]byte, robot Pos, instr int) Pos {
	// Check what is the last position that the robot can move
	ok, seek_pos := seekMove(warehouse, robot, instr)
	if !ok || (robot.x == seek_pos.x && robot.y == seek_pos.y) {
		fmt.Println("Can't move", robot, instr)
		return robot
	}

	// Move robot (and boxes)
	// fmt.Println("Move boxes from", robot, "to", seek_pos)
	return rotateMove(warehouse, robot, seek_pos, instr)
}

func main() {
	input := test_input
	// input = test_input2
	input = ex_input
	input = real_input
	warehouse, robot, instructions := parseInput(input)

	result := 0
	for _, instr := range instructions {
		robot = checkMove(warehouse, robot, instr)
		fmt.Println(robot)
		result = getSum(warehouse)
	}

	fmt.Println(result)
}
