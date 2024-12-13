package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

type Claw struct {
	x int
	y int
}

var r = regexp.MustCompile(`.*: X.(\d+), Y.(\d+)`)

func parseClaw(input string) Claw {
	all_groups := r.FindStringSubmatch(input)
	this_x, _ := strconv.Atoi(all_groups[1])
	this_y, _ := strconv.Atoi(all_groups[2])
	return Claw{this_x, this_y}
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	// fmt.Println(input)

	result := 0
	for i := 0; i < len(input); i += 4 {
		// Parse
		a := parseClaw(input[i])
		b := parseClaw(input[i+1])
		target := parseClaw(input[i+2])
		target = Claw{target.x + 10000000000000, target.y + 10000000000000}

		// It's math
		// 8400 = a * 94 + b * 22
		// a * 94 = 8400 - 22 * b
		// a = ( 8400 - 22 * b ) / 94
		//
		// 5400 = a * 34 + b * 67
		// 5400 =  ( 8400 - 22 * b ) * 34 / 94 + b * 67
		// 5400 * 94 = 8400 * 34 - 34 * 22 * b + b * 67 * 94
		// b * ( 67 * 94 - 34 * 22 ) = 5400 * 94 - 8400 * 34

		num_a, num_b := 0, 0
		// Calc
		left := target.y*a.x - target.x*a.y
		right := b.y*a.x - b.x*a.y
		// fmt.Println(left, right)
		if left%right == 0 {
			num_b = left / right
		} else {
			fmt.Println("No solution")
			continue
		}
		// fmt.Println(num_b)
		left = target.x - num_b*b.x
		right = a.x
		// fmt.Println(left, right)
		if left%right == 0 {
			num_a = left / right
		} else {
			fmt.Println("No solution")
			continue
		}

		fmt.Println(num_a, num_b)
		result += num_a*3 + num_b

		// break
	}
	fmt.Println(result)
}
