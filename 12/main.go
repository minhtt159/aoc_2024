package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
)

func read_input(file_name string) []string {
	result := []string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

var (
	dir_x = []int{-1, 0, 0, 1}
	dir_y = []int{0, -1, 1, 0}
)

type Plot struct {
	id      rune
	area    int
	peri    int
	corners int
}

func BFS(garden []string, visited map[int]map[int]bool, target rune, x int, y int) (bool, Plot) {
	if garden[x][y] != byte(target) {
		return false, Plot{0, 0, 0, 0}
	}
	if visited[x][y] {
		return true, Plot{0, 0, 0, 0}
	}
	visited[x][y] = true
	totalArea := 1
	totalPeri := 0

	// for part 2
	totalCorners := 0
	adj := []int{}

	for k := 0; k < 4; k++ {
		new_x := x + dir_x[k]
		new_y := y + dir_y[k]
		if new_x < 0 || new_x >= len(garden) || new_y < 0 || new_y >= len(garden[0]) {
			// Out of bound
			totalPeri++
		} else if ok, nextPlot := BFS(garden, visited, target, new_x, new_y); ok {
			// Neighbor plot is the same kind as this plot
			totalArea += nextPlot.area
			totalPeri += nextPlot.peri

			// for part 2
			totalCorners += nextPlot.corners
			adj = append(adj, k)
		} else {
			// Neighbor plot is not the same kind as this plot
			totalPeri++
		}
	}

	// for part 2
	sort.Ints(adj)
	if len(adj) == 0 {
		fmt.Println("Plot", x, y, "has no adjecent plot")
		totalCorners += 4
	} else if len(adj) == 1 {
		// True corner
		fmt.Println("Plot", x, y, "has only one adjecent plot")
		totalCorners += 3
	} else if reflect.DeepEqual(adj, []int{0, 2}) || reflect.DeepEqual(adj, []int{1, 3}) {
		// L shape
		fmt.Println("Plot", x, y, "has L shape")
		totalCorners += 2
	}

	return true, Plot{target, totalArea, totalPeri, totalCorners}
}

func main() {
	file_name := "input.txt"
	// file_name := "real_input.txt"
	input := read_input(file_name)
	input = []string{"AAAA", "BBCD", "BBCC", "EEEC"}
	// input = []string{"OOOOO", "OXOXO", "OOOOO", "OXOXO", "OOOOO"}

	visited := make(map[int]map[int]bool)
	for i, line := range input {
		visited[i] = make(map[int]bool, len(line))
		for j := range line {
			visited[i][j] = false
		}
	}

	allPlots := []Plot{}
	for i, line := range input {
		for j, c := range line {
			if visited[i][j] {
				continue
			}
			_, plot := BFS(input, visited, c, i, j)
			allPlots = append(allPlots, plot)
		}
	}

	result := 0
	for _, plot := range allPlots {
		fmt.Printf("Plot %c has area %d, perimeter %d, and corner %d\n", plot.id, plot.area, plot.peri, plot.corners)
		result += plot.area * plot.peri
	}
	fmt.Println(result)
}
