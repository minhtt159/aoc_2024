package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	id   rune
	area int
	peri int
}

func BFS(garden []string, visited map[int]map[int]bool, target rune, x int, y int) (bool, Plot) {
	if garden[x][y] != byte(target) {
		return false, Plot{0, 0, 0}
	}
	if visited[x][y] {
		return true, Plot{0, 0, 0}
	}
	visited[x][y] = true
	totalArea := 1
	totalPeri := 0
	for k := 0; k < 4; k++ {
		new_x := x + dir_x[k]
		new_y := y + dir_y[k]
		if new_x < 0 || new_x >= len(garden) || new_y < 0 || new_y >= len(garden[0]) {
			// out of bound
			totalPeri++
			continue
		}
		if ok, nextPlot := BFS(garden, visited, target, new_x, new_y); ok {
			// add perimeter and area
			totalArea += nextPlot.area
			// look at how many sides are connected
			for v := 0; v < 4; v++ {
				next_x := new_x + dir_x[v]
				next_y := new_y + dir_y[v]
				if next_x < 0 || next_x >= len(garden) || next_y < 0 || next_y >= len(garden[0]) {
					// out of bound, add perimeter
					totalPeri++
				} else if garden[next_x][next_y] != byte(target) {
					// not the same plot, add perimeter
					totalPeri++
				} else if garden[next_x][next_y] == byte(target) {
					// same plot, remove perimeter
					totalPeri--
				}
			}
			// inherit perimeter
			totalPeri += nextPlot.peri
			fmt.Println("This plot", x, y, "need", totalPeri, "fences")
		}
	}

	fmt.Println("Visited", string(target), x, y, totalArea, totalPeri)
	return true, Plot{target, totalArea, totalPeri}
}

func main() {
	file_name := "input.txt"
	input := read_input(file_name)
	input = []string{"AAAA", "BBCD", "BBCC", "EEEC"}

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

	for _, plot := range allPlots {
		log.Printf("Plot %c has area %d and perimeter %d", plot.id, plot.area, plot.peri)
	}
}
