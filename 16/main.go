package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var example_input string // 7036, 45

var input1 = `######
#...E#
#....#
#....#
#S...#
######
` // 1006, 7

//go:embed input2.txt
var input2 string // 11048, 64

//go:embed real_input.txt
var real_input string

type Pos struct {
	x int
	y int
	d int
}

type Route struct {
	steps []Pos
	pos   Pos
	score int
}

var (
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, -1, 0, 1}
)

const (
	U = iota
	R
	D
	L
)

func parseMaze(input string) ([][]byte, Pos, Pos) {
	all_lines := strings.Split(input, "\n")
	maze := make([][]byte, len(all_lines)-1)
	pos_s := Pos{0, 0, L}
	pos_e := Pos{0, 0, -1}
	for i, line := range all_lines {
		if len(line) == 0 {
			continue
		}
		maze[i] = make([]byte, len(line))
		for j, c := range line {
			maze[i][j] = byte(c)
			switch maze[i][j] {
			case 'S':
				pos_s = Pos{i, j, L}
			case 'E':
				pos_e = Pos{i, j, -1}
			}
		}
	}
	return maze, pos_s, pos_e
}

func debugMaze(maze [][]byte, route Route) {
	fmt.Println("Score :", route.score)

	getStep := func(x int, y int) (byte, bool) {
		for i, step := range route.steps {
			if step.x == x && step.y == y {
				switch i {
				case 0:
					return 'S', true
				case len(route.steps) - 1:
					return 'E', true
				default:
					// return byte(step.d + '0'), true
					switch step.d {
					case 0:
						return '^', true
					case 1:
						return '<', true
					case 2:
						return 'v', true
					case 3:
						return '>', true
					}
				}
			}
		}
		return 0, false
	}

	for i := 0; i < len(maze); i++ {
		line := make([]byte, len(maze))
		copy(line, maze[i])
		for j := 0; j < len(maze[0]); j++ {
			if r, ok := getStep(i, j); ok {
				line[j] = r
			}
		}
		fmt.Println(string(line))
	}
}

func getNextPos(cur Pos) []Pos {
	next_cur := [4]Pos{}
	for i := 0; i < 4; i++ {
		next_cur[i] = Pos{cur.x + dir_x[i], cur.y + dir_y[i], i}
	}
	// Usually, we'll get all 4 positions
	// But in this case, it's stupid
	switch cur.d {
	case U:
		return []Pos{next_cur[U], next_cur[L], next_cur[R]}
	case R:
		return []Pos{next_cur[U], next_cur[R], next_cur[D]}
	case D:
		return []Pos{next_cur[D], next_cur[L], next_cur[R]}
	case L:
		return []Pos{next_cur[U], next_cur[L], next_cur[D]}
	default:
		log.Fatal("Unknown direction")
	}
	return []Pos{}
}

func Dijkstra(maze [][]byte, cur Pos) (int, int) {
	visited := make(map[Pos]int) // dijkstra visited map

	bestScore := math.MaxInt32       // infitity
	bestRoute := make(map[int][]Pos) // all the steps that a part of the smallest route

	queue := []Route{
		{
			[]Pos{cur}, // steps in this route
			cur,        // current position
			0,          // score of this route
		},
	}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.steps) > len(maze)*len(maze[0]) || state.score > bestScore {
			// That'll be too much
			continue
		}

		if maze[state.pos.x][state.pos.y] == 'E' {
			// Done
			// debugMaze(maze, state)
			if state.score <= bestScore {
				// Smallest score found
				bestScore = state.score
				// bestRoute map
				bestRoute[state.score] = append(bestRoute[state.score], state.steps...)
				continue
			}
		}

		for _, next_pos := range getNextPos(state.pos) {
			if maze[next_pos.x][next_pos.y] != '#' {
				// Move forward
				this_score := state.score + 1
				if next_pos.d != state.pos.d {
					// Add rotate cost
					this_score += 1000
				}

				if previous_score, ok := visited[next_pos]; ok {
					if previous_score < this_score {
						// We met this position before
						// but this_score is higher than previous_score
						continue
					}
				}
				// Commit this_score because it's the best score
				visited[next_pos] = this_score

				// Make a copy of the steps
				newSteps := make([]Pos, len(state.steps))
				copy(newSteps, state.steps)

				// Queue this position
				queue = append(queue, Route{
					append(newSteps, next_pos),
					next_pos,
					this_score,
				})
			}
		}
	}

	buffer := make(map[Pos]int) // all steps that make the best route
	for _, v := range bestRoute[bestScore] {
		buffer[Pos{v.x, v.y, -1}]++
	}
	// fmt.Println(buffer)

	debugMaze(maze, Route{
		bestRoute[bestScore],
		cur,
		bestScore,
	})
	return bestScore, len(buffer)
}

func main() {
	input := example_input
	// input = input1
	input = real_input
	maze, start, _ := parseMaze(input)
	debugMaze(maze, Route{[]Pos{}, start, 0})

	score, steps := Dijkstra(maze, start)
	fmt.Println(score, steps)
}
