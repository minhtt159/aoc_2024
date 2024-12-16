package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var test_input string // 1930, 1206

var (
	input_2 = "AAAA\nBBCD\nBBCC\nEEEC"            // 140, 80
	input_3 = "OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO" // 772, 436
)

//go:embed real_input.txt
var real_input string

func read_input(input string) []string {
	result := []string{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		result = append(result, line)
	}
	return result
}

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

type Plot struct {
	x int
	y int
	d int
}

func getNeighbors(target Plot) []Plot {
	return []Plot{
		{target.x + dir_x[U], target.y + dir_y[U], U},
		{target.x + dir_x[L], target.y + dir_y[L], L},
		{target.x + dir_x[R], target.y + dir_y[R], R},
		{target.x + dir_x[D], target.y + dir_y[D], D},
	}
}

func DFS(garden []string, visited [][]bool, target Plot, plots *[]Plot) (int, int) {
	visited[target.x][target.y] = true
	totalArea := 1
	totalPeri := 0

	for _, plot := range getNeighbors(target) {
		if plot.x < 0 || plot.x >= len(garden) || plot.y < 0 || plot.y >= len(garden[0]) ||
			garden[target.x][target.y] != garden[plot.x][plot.y] {
			// Out of bound or not the same kind
			totalPeri++

			// For part 2,
			// Add the neighbor plot as the inner/outer bounds
			*plots = append(*plots, plot)

		} else if !visited[plot.x][plot.y] {
			// Visit the neighbor and add the area and perimeter
			a, p := DFS(garden, visited, plot, plots)
			totalArea += a
			totalPeri += p
		}
	}

	return totalArea, totalPeri
}

func countCorner(plots []Plot) int {
	listSides := make(map[Plot]bool)
	totalSide := 0

	// Sort the plots by row then column,
	// because when iterating the list,
	// the sides are added sequentially
	sort.Slice(plots, func(i int, j int) bool {
		if plots[i].x == plots[j].x {
			return plots[i].y < plots[j].y
		}
		return plots[i].x < plots[j].x
	})
	// fmt.Println("Plots", plots)

	// Plots is a list of outer/inner bounds
	for _, plot := range plots {
		found := false

		// check if the neighbor with this direction is already in the list
		for _, neighbor := range getNeighbors(plot) {
			neighbor.d = plot.d
			if _, ok := listSides[neighbor]; ok {
				found = true
				// fmt.Println("Found side", plot, neighbor)
				break
			}
		}

		// if this side hasn't been added before
		if !found {
			totalSide++
		}
		listSides[plot] = true
	}

	// fmt.Println(listSides)
	return totalSide
}

func main() {
	file := test_input
	// file = input_2
	// file = input_3
	file = real_input
	input := read_input(file)
	for _, line := range input {
		fmt.Println(line)
	}

	visited := make([][]bool, len(input))
	for i, line := range input {
		visited[i] = make([]bool, len(line))
		for j := range line {
			visited[i][j] = false
		}
	}

	r1 := 0
	r2 := 0
	for i, line := range input {
		for j := range line {
			if !visited[i][j] {
				plots := make([]Plot, 0)
				area, peri := DFS(input, visited, Plot{i, j, 0}, &plots)
				corners := countCorner(plots)
				fmt.Printf("Plot %c has area %d, perimeter %d, and corner %d\n",
					input[i][j],
					area,
					peri,
					corners,
				)
				r1 += area * peri
				r2 += area * corners
				// log.Fatal("break")
			}
		}
	}

	fmt.Println("Result 1:", r1)
	fmt.Println("Result 2:", r2)
}
