package main

import (
	"log"
	"math"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Cache struct {
	map_type int
	start    Pos
	end      Pos
}

type Route struct {
	path   []Pos  // path from a -> b
	dir    []int  // representation in UDLR
	output string // output string
	cur    Pos    // helper for dijkstra
	score  int    // helper fot dijkstra
}

var cache = make(map[Cache][]string)

// Find the fastest path from start to end
// and cache them
func Dijkstra(input [][3]int, start Pos, end Pos) []string {
	// default to a direction pad
	map_type := 0
	if len(input) == 4 {
		// number pad
		map_type = 1
	}

	allRoute := []string{}

	// If start and end are the same, still cache it
	if start.x == end.x && start.y == end.y {
		allRoute = []string{"A"}
		cache[Cache{map_type, start, end}] = allRoute
		return allRoute
	}

	// Begin Dijkstra
	minScore := math.MaxInt64
	visited := make(map[Pos]bool)
	queue := []Route{
		{
			[]Pos{start},
			[]int{},
			"",
			start,
			0,
		},
	}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		// This path is already longer than the shortest path
		if state.score > minScore {
			continue
		}
		cur_pos := state.cur
		// End condition
		if cur_pos.x == end.x && cur_pos.y == end.y {
			if state.score <= minScore {
				// Shortest path found
				state.output += "A"
				if state.score < minScore {
					// fmt.Println("Found new shortest path", state)
					minScore = state.score
					allRoute = []string{}
				}
				allRoute = append(allRoute, state.output)
			}
			continue
		}
		// visited
		if _, ok := visited[cur_pos]; ok {
			continue
		}
		// if not visited, queue the neighbor
		for i := 0; i < 4; i++ {
			next_pos := Pos{cur_pos.x + dir_x[i], cur_pos.y + dir_y[i]}
			// out of map
			if next_pos.x < 0 || next_pos.x >= len(input) || next_pos.y < 0 || next_pos.y >= len(input[0]) {
				continue
			}
			// useless neighbor
			if input[next_pos.x][next_pos.y] == N {
				continue
			}
			// next_pos score
			next_score := state.score + 1
			// it's always expensive to change the direction
			if len(state.dir) > 0 && state.dir[len(state.dir)-1] != i {
				next_score++
			}
			// next_pos path
			next_path := make([]Pos, len(state.path))
			copy(next_path, state.path)
			// next_pos dir
			next_dir := make([]int, len(state.dir))
			copy(next_dir, state.dir)
			// next_pos output
			next_output := strings.Clone(state.output)
			switch i {
			case U:
				next_output += "^"
			case L:
				next_output += "<"
			case R:
				next_output += ">"
			case D:
				next_output += "v"
			default:
				log.Fatal("Invalid direction")
			}
			// queue the next_pos
			queue = append(queue, Route{
				append(next_path, next_pos),
				append(next_dir, i),
				next_output,
				next_pos,
				next_score,
			})
		}
	}

	// allRoute is a slices of shortest path outputs
	cache[Cache{map_type, start, end}] = allRoute
	return allRoute
}
