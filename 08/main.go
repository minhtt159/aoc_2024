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

type Antenna struct {
	x int
	y int
	d rune // I don't need this field, but it's here for consistency
}

func fillx(input []string, row int, col int, c rune) {
	input[row] = input[row][:col] + string(c) + input[row][col+1:]
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)

	allAntenna := make(map[rune][]Antenna)

	for x, input := range input {
		for y, rune := range input {
			if rune == '.' {
				continue
			}
			thisType := allAntenna[rune]
			thisAntenna := Antenna{
				x,
				y,
				rune,
			}
			if len(thisType) == 0 {
				allAntenna[rune] = []Antenna{thisAntenna}
			} else {
				allAntenna[rune] = append(allAntenna[rune], thisAntenna)
			}
		}
	}

	antinodes := []Antenna{}
	for _, thisType := range allAntenna {
		// fmt.Println(string(r), thisType)

		if len(thisType) <= 1 {
			continue
		}
		for i := 0; i < len(thisType); i++ {
			antenna_i := thisType[i]
			for j := i + 1; j < len(thisType); j++ {
				antenna_j := thisType[j]

				// Distance between i and j
				d_x := antenna_i.x - antenna_j.x
				d_y := antenna_i.y - antenna_j.y
				// fmt.Println(antenna_i, antenna_j, d_x, d_y)

				lapse := 1000 // Just an arbitrary number, phase 1 = 1, phase 2 = 1000, etc
				// Potential antinode
				for k := 0; k < lapse; k++ {
					a_x := antenna_i.x + d_x*k
					a_y := antenna_i.y + d_y*k
					// Check if a and b are within the grid
					if a_x >= 0 && a_x < len(input) && a_y >= 0 && a_y < len(input[0]) {
						if input[a_x][a_y] == '.' {
							fillx(input, a_x, a_y, '#')
						} else {
							// fmt.Println("Potential antinode a is not empty")
						}
						antinode := Antenna{a_x, a_y, '#'}
						if !slices.Contains(antinodes, antinode) {
							antinodes = append(antinodes, antinode)
						}
					} else {
						break
					}
				}
				for k := 0; k < lapse; k++ {
					b_x := antenna_j.x - d_x*k
					b_y := antenna_j.y - d_y*k
					if b_x >= 0 && b_x < len(input) && b_y >= 0 && b_y < len(input[0]) {
						if input[b_x][b_y] == '.' {
							fillx(input, b_x, b_y, '#')
						} else {
							// fmt.Println("Potential antinode b is not empty")
						}
						antinode := Antenna{b_x, b_y, '#'}
						if !slices.Contains(antinodes, antinode) {
							antinodes = append(antinodes, antinode)
						}
					} else {
						break
					}
				}
			}
		}
	}

	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println(len(antinodes))
}
