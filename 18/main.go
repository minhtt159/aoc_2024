package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
)

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

var (
	r            = `(\d+),(\d+)`
	max_x, max_y = 71, 71
	// max_x, max_y = 7, 7
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, -1, 0, 1}
)

type Pos struct {
	x int
	y int
	d int
}

type Route struct {
	path  []Pos
	cur   Pos
	score int
}

const (
	U = iota
	L
	D
	R
)

func parseInput(input string) []Pos {
	reg := regexp.MustCompile(r)
	result := []Pos{}
	parsedLine := reg.FindAllStringSubmatch(input, -1)
	for _, pair := range parsedLine {
		y, _ := strconv.Atoi(pair[1])
		x, _ := strconv.Atoi(pair[2])
		result = append(result, Pos{x, y, -1})
	}
	return result
}

func dijkstra(cur Pos, obs []Pos) ([]Pos, int) {
	visited := make(map[Pos]int)
	bestPath := []Pos{}
	bestScore := math.MaxInt

	queue := []Route{
		{
			[]Pos{cur},
			cur,
			0,
		},
	}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if len(top.path) > max_x*max_y || top.score > bestScore {
			// too big
			continue
		}

		if top.cur.x == max_x-1 && top.cur.y == max_y-1 {
			// fmt.Println("Found", top.path, top.score)
			if top.score < bestScore {
				bestScore = top.score
				bestPath = top.path
			}
			break
		}

		for i := 0; i < 4; i++ {
			next := Pos{top.cur.x + dir_x[i], top.cur.y + dir_y[i], -1}
			if !slices.Contains(obs, next) &&
				top.cur.x >= 0 && top.cur.x < max_x &&
				top.cur.y >= 0 && top.cur.y < max_y {
				this_score := top.score + 1

				if prev_score, ok := visited[next]; ok {
					if prev_score < this_score {
						continue
					}
				}

				visited[next] = top.score

				newPath := make([]Pos, len(top.path))
				copy(newPath, top.path)

				queue = append(queue, Route{
					append(newPath, next),
					next,
					this_score,
				})
			}
		}
	}
	return bestPath, bestScore
}

func main() {
	input_str := ex_input
	input_str = real_input
	obs := parseInput(input_str)
	// fmt.Println(len(obs))

	// for i := 0; i < max_x; i++ {
	// 	for j := 0; j < max_y; j++ {
	// 		if slices.Contains(obs, Pos{i, j, -1}) {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }
	num_obs := 1024
	_, score := dijkstra(Pos{0, 0, -1}, obs[:num_obs])
	fmt.Println(score)

	for i := num_obs; i < len(obs); i++ {
		path, _ := dijkstra(Pos{0, 0, -1}, obs[:i])

		if len(path) == 0 {
			fmt.Println("p2", i-1, obs[i-1])
			break
		}
	}
}
