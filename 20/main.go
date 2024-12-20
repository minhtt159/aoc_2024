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

type Route struct {
	path  []Pos
	wall1 Pos
	wall2 Pos
	cur   Pos
	score int
}

var (
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, -1, 0, 1}
)

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
	fmt.Println("Score :", route.score)

	wallList := []Pos{route.wall1, route.wall2}
	for i := 0; i < len(maze); i++ {
		line := make([]byte, len(maze))
		copy(line, maze[i])
		for j := 0; j < len(maze[0]); j++ {
			if slices.Contains(wallList, Pos{i, j, -1}) {
				line[j] = 'X'
			} else if slices.Contains(route.path, Pos{i, j, -1}) {
				line[j] = '0'
			}
		}
		fmt.Println(string(line))
	}
}

var nonBreakScore int

func Dijkstra(maze [][]byte, cur Pos, wallBreak bool) (int, int) {
	bestScore := 0
	bestRoutes := []Route{}

	S := Route{
		[]Pos{cur}, // steps in this route
		Pos{-1, -1, -1},
		Pos{-1, -1, -1},
		cur,
		0, // score of this route
	}

	if !wallBreak {
		S.wall1 = Pos{len(maze), -1, -1}
		S.wall2 = Pos{len(maze), -1, -1}
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
			debugMaze(maze, state)
			bestScore = state.score
			bestRoutes = append(bestRoutes, state)
			if !wallBreak {
				break
			}
			continue
		}

		for i := 0; i < 4; i++ {
			next_pos := Pos{state.cur.x + dir_x[i], state.cur.y + dir_y[i], -1}
			if next_pos.x < 0 || next_pos.x >= len(maze) || next_pos.y < 0 || next_pos.y >= len(maze[0]) {
				// out of bound
				continue
			}
			if (state.wall1.x == -1 && state.wall2.x == -1) || !wallBreak {

				if slices.Contains(state.path, next_pos) {
					// Already visited
					continue
				}

				if maze[next_pos.x][next_pos.y] == '#' && wallBreak {
					if state.wall1.x == 1 {
						state.wall1 = next_pos
					} else if state.wall1.x != 1 && state.wall2.x == -1 {
						state.wall2 = next_pos
					} else {
						continue
					}
				}

				// Make a copy of the steps
				newSteps := make([]Pos, len(state.path))
				copy(newSteps, state.path)

				// Queue this position
				queue = append(queue, Route{
					append(newSteps, next_pos),
					Pos{state.wall1.x, state.wall1.y, -1},
					Pos{state.wall2.x, state.wall2.y, -1},
					next_pos,
					state.score + 1,
				})
			}
		}
	}

	fmt.Println("Best routes:", bestRoutes)
	return bestScore, len(bestRoutes)
}

func main() {
	input := ex_input
	maze, start := parseMaze(input)

	r, s := Dijkstra(maze, start, false)
	fmt.Println("Result:", r, s)

	r2, s2 := Dijkstra(maze, start, true)
	fmt.Println("Result:", r2, s2)
}
