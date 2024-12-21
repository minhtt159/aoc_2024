package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var ex_input string

// 029A <- 29 * 68
// 980A <- 980 * 60
// 179A <- 179 * 68
// 456A <- 456 * 64
// 379A <- 379 * 64

//go:embed real_input.txt
var real_input string

const (
	U = iota
	L
	D
	R
	N = -1
	A = 10
)

type Pos struct {
	x int
	y int
}

type Cache struct {
	map_type int // num or dir
	start    Pos
	end      Pos
}

type Route struct {
	path   []Pos
	dir    []int
	output string
	cur    Pos
	score  int
}

type Case struct {
	input string
	round int
}

var (
	// num pad
	num_map = [][3]int{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
		{N, 0, A},
	}
	// direction pad
	dir_map = [][3]int{
		{N, U, A},
		{L, D, R},
	}
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, -1, 0, 1}
	// route cache
	cache = make(map[Cache][]Route)
	case1 = Case{
		ex_input,
		2,
	}
	case2 = Case{
		ex_input,
		25,
	}
	case3 = Case{
		real_input,
		2,
	}
	case4 = Case{
		real_input,
		25,
	}
)

// Find the fastest path
// from start to end
// and cache them
func Dijkstra(input [][3]int, start Pos, end Pos) []Route {
	allRoute := []Route{}
	minScore := math.MaxInt64
	// result := Route{
	// 	[]Pos{},
	// 	[]int{},
	// 	"",
	// 	start,
	// 	math.MaxInt64, // crazy high number
	// }
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

		if state.score > minScore {
			continue
		}
		cur_pos := state.cur

		if cur_pos.x == end.x && cur_pos.y == end.y {
			if state.score <= minScore {
				state.output += "A"
				if state.score < minScore {
					// fmt.Println("Found new shortest path", state)
					minScore = state.score
					allRoute = []Route{}
				}
				allRoute = append(allRoute, state)
			}
			continue
		}

		if _, ok := visited[cur_pos]; ok {
			// visited
			continue
		}

		for i := 0; i < 4; i++ {
			next_score := state.score + 1
			next_pos := Pos{cur_pos.x + dir_x[i], cur_pos.y + dir_y[i]}
			// it's always expensive to change the direction
			if len(state.dir) > 0 && state.dir[len(state.dir)-1] != i {
				next_score++
			}

			if next_pos.x < 0 || next_pos.x >= len(input) || next_pos.y < 0 || next_pos.y >= len(input[0]) {
				// out of map
				continue
			}
			if input[next_pos.x][next_pos.y] == N {
				continue
			}

			// copy next path
			next_path := make([]Pos, len(state.path))
			copy(next_path, state.path)
			// copy next dir
			next_dir := make([]int, len(state.dir))
			copy(next_dir, state.dir)
			// copy next output
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

			queue = append(queue, Route{
				append(next_path, next_pos),
				append(next_dir, i),
				next_output,
				next_pos,
				next_score,
			})
		}
	}

	// default to a direction pad
	map_type := 0
	if len(input) == 4 {
		// number pad
		map_type = 1
	}
	// fmt.Println("Caching", map_type, input[start.x][start.y], input[end.x][end.y], len(allRoute))
	// for _, route := range allRoute {
	// 	fmt.Println(route.output)
	// }
	cache[Cache{map_type, start, end}] = allRoute
	return allRoute
}

func findPosInMap(input [][3]int, r rune) Pos {
	var n1 int
	switch r {
	case 'A':
		n1 = A
	case 'N':
		n1 = N
	case '^':
		n1 = U
	case '<':
		n1 = L
	case 'v':
		n1 = D
	case '>':
		n1 = R
	default:
		n1, _ = strconv.Atoi(string(r))
	}
	for i, row := range input {
		for j, n := range row {
			if n == n1 {
				return Pos{i, j}
			}
		}
	}
	return Pos{-1, -1}
}

func main() {
	// Init cache
	for i, line1 := range num_map {
		for j, r1 := range line1 {
			for k, line2 := range num_map {
				for v, r2 := range line2 {
					if i == k && j == v {
						cache[Cache{1, Pos{i, j}, Pos{k, v}}] = []Route{{output: "A"}}
						continue
					}
					if r1 == N || r2 == N {
						// useless
						continue
					}
					Dijkstra(num_map, Pos{i, j}, Pos{k, v})
				}
			}
		}
	}
	for i, line1 := range dir_map {
		for j, r1 := range line1 {
			for k, line2 := range dir_map {
				for v, r2 := range line2 {
					if i == k && j == v {
						cache[Cache{0, Pos{i, j}, Pos{k, v}}] = []Route{{output: "A"}}
						continue
					}
					if r1 == N || r2 == N {
						// useless
						continue
					}
					Dijkstra(dir_map, Pos{i, j}, Pos{k, v})
				}
			}
		}
	}

	xxx := case2
	input, round := xxx.input, xxx.round
	combos := strings.Split(strings.Trim(input, "\n"), "\n")

	sum := 0
	for _, combo := range combos {
		weight, _ := strconv.Atoi(strings.Trim(combo, "A"))

		//
		stage_1 := []string{}
		start := Pos{3, 2} // A of num map
		for _, r := range combo {
			end := findPosInMap(num_map, r)

			possible_path := cache[Cache{1, start, end}]
			next_step := []string{}

			for _, new_path := range possible_path {
				if len(stage_1) == 0 {
					next_step = append(next_step, new_path.output)
				} else {
					for _, path := range stage_1 {
						next_step = append(next_step, path+new_path.output)
					}
				}
			}

			start = end
			stage_1 = next_step
		}
		fmt.Println("Path", combo, "=>")
		for _, path := range stage_1 {
			fmt.Println(path, len(path))
		}
		// <A^A>^^AvvvA 12

		last_round := stage_1
		for i := 0; i < round; i++ {
			// 25 rounds
			this_round := []string{}
			start = Pos{0, 2} // A of dir map
			for _, stage := range last_round {
				this_step := []string{}
				for _, r := range stage {
					end := findPosInMap(dir_map, r)

					possible_path := cache[Cache{0, start, end}]
					next_step := []string{}

					for _, new_path := range possible_path {
						if len(this_step) == 0 {
							next_step = append(next_step, new_path.output)
						} else {
							for _, path := range this_step {
								next_step = append(next_step, path+new_path.output)
							}
						}
					}

					start = end
					this_step = next_step
				}
				// fmt.Println("Path", stage, "=>", this_step, len(this_step))
				this_round = append(this_round, this_step...)
			}
			fmt.Println("Path", last_round, "=>")
			// for _, path := range this_round {
			// 	fmt.Println(path, len(path))
			// }
			last_round = []string{this_round[0]}
			fmt.Println(this_round[0], len(this_round[0]))
			// v<<A>>^A<A>AvA<^AA>A<vAAA>^A 28
		}

		last_length := math.MaxInt64

		for _, path := range last_round {
			// fmt.Println(path, len(path))
			if len(path) < last_length {
				last_length = len(path)
			}
		}
		// <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A 68

		sum += last_length * weight
		// break
	}

	fmt.Println("Sum", sum)
}
