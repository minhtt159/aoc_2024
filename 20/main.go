package main

import (
	_ "embed"
	"fmt"
	"math"
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

type Route struct {
	path []Pos
	cur  Pos
}

const (
	U = iota
	L
	D
	R
)

func parseMaze(input string) ([][]byte, Pos) {
	all_lines := strings.Split(input, "\n")
	maze := make([][]byte, len(all_lines)-1)
	pos_s := Pos{0, 0, -1}
	for i, line := range all_lines {
		if len(line) == 0 {
			continue
		}
		maze[i] = make([]byte, len(line))
		for j, c := range line {
			maze[i][j] = byte(c)
			switch maze[i][j] {
			case 'S':
				pos_s = Pos{i, j, -1}
				// case 'E':
				// 	pos_e = Pos{i, j, -1}
			}
		}
	}
	return maze, pos_s
}

func debugMaze(maze [][]byte, route Route) {
	fmt.Println("Score :", len(route.path))

	for i := 0; i < len(maze); i++ {
		line := make([]byte, len(maze))
		copy(line, maze[i])
		for j := 0; j < len(maze[0]); j++ {
			if slices.Contains(route.path, Pos{i, j, -1}) {
				line[j] = '0'
			}
		}
		fmt.Println(string(line))
	}
}

type Case struct {
	input     string
	target    int
	cheat_num int
}

var (
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, -1, 0, 1}
	case1 = Case{
		ex_input,
		2,
		2,
	}
	case2 = Case{
		ex_input,
		50,
		20,
	}
	case3 = Case{
		real_input,
		100,
		2,
	}
	case4 = Case{
		real_input,
		100,
		20,
	}
)

func Dijkstra(maze [][]byte, cur Pos) (int, []Pos) {
	bestScore := 0
	bestRoute := []Pos{}
	S := Route{
		[]Pos{cur}, // steps in this route
		cur,
	}

	queue := []Route{S}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.path) > len(maze)*len(maze[0]) {
			// That'll be too much
			continue
		}

		if maze[state.cur.x][state.cur.y] == 'E' {
			// Done
			score := len(state.path) - 1
			// debugMaze(maze, state)
			bestScore = score
			bestRoute = state.path
			break
		}

		for i := 0; i < 4; i++ {
			next_pos := Pos{state.cur.x + dir_x[i], state.cur.y + dir_y[i], -1}
			if next_pos.x < 0 || next_pos.x >= len(maze) || next_pos.y < 0 || next_pos.y >= len(maze[0]) {
				// out of bound
				continue
			}
			if maze[next_pos.x][next_pos.y] != '#' {
				if slices.Contains(state.path, next_pos) {
					// Already visited
					continue
				}

				// Make a copy of the steps
				newSteps := make([]Pos, len(state.path))
				copy(newSteps, state.path)
				newRoute := Route{
					append(newSteps, next_pos),
					next_pos,
				}

				// Queue this position
				queue = append(queue, newRoute)
			}
		}
	}

	return bestScore, bestRoute
}

func main() {
	xxx := case4
	input, target, cheat_num := xxx.input, xxx.target, xxx.cheat_num
	// input := real_input
	maze, start := parseMaze(input)

	max_score, route := Dijkstra(maze, start)
	fmt.Println("Origin path:", max_score)

	r := 0
	for a := 0; a < len(route); a++ {
		pa := route[a]
		for b := a + target + 1; b < len(route); b++ {
			pb := route[b]
			dx := math.Abs(float64(pb.x - pa.x))
			dy := math.Abs(float64(pb.y - pa.y))
			d := int(dx + dy)

			if d <= cheat_num && b-a-d >= target {
				// fmt.Println("Cheat:", a, b, d)
				r++
			}

		}
	}
	fmt.Println("Cheat:", r)
}
