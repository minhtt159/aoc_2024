package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

type Node struct {
	x int
	y int
}

var (
	dir_x = []int{-1, 0, 0, 1}
	dir_y = []int{0, -1, 1, 0}
)

func BFS(graph []string, path []Node) []Node {
	lastNode := path[len(path)-1]
	lastRune := graph[lastNode.x][lastNode.y]
	if lastRune == '0' {
		return []Node{lastNode}
	}
	var result []Node
	for i := 0; i < 4; i++ {
		new_x := lastNode.x + dir_x[i]
		new_y := lastNode.y + dir_y[i]
		if new_x < 0 || new_x >= len(graph) || new_y < 0 || new_y >= len(graph[0]) {
			// out of map
			continue
		}
		if graph[new_x][new_y] != lastRune-1 {
			// not a trail
			continue
		}
		newPath := append(path, Node{new_x, new_y})
		thisResult := BFS(graph, newPath)
		for _, iResult := range thisResult {
			if slices.Contains(result, iResult) {
				continue
			}
			result = append(result, iResult)
		}
	}
	return result
}

func BFS_2(graph []string, path []Node) int {
	lastNode := path[len(path)-1]
	lastRune := graph[lastNode.x][lastNode.y]
	if lastRune == '0' {
		return 1
	}
	result := 0
	for i := 0; i < 4; i++ {
		new_x := lastNode.x + dir_x[i]
		new_y := lastNode.y + dir_y[i]
		if new_x < 0 || new_x >= len(graph) || new_y < 0 || new_y >= len(graph[0]) {
			// out of map
			continue
		}
		if graph[new_x][new_y] != lastRune-1 {
			// not a trail
			continue
		}
		newPath := append(path, Node{new_x, new_y})
		result += BFS_2(graph, newPath)
	}
	return result
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	fmt.Println(input)

	total := 0
	total_2 := 0
	for i, line := range input {
		for j, thisRune := range line {
			if thisRune == '9' {
				path := []Node{{i, j}}
				thisTrail := BFS(input, path)
				fmt.Println(thisTrail)
				total += len(thisTrail)

				total_2 += BFS_2(input, path)
			}
		}
	}
	fmt.Println(total)
	fmt.Println(total_2)
}
