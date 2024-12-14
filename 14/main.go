package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	r         = regexp.MustCompile(`p=(\d+),(\d+) v=(.*),(.*)`)
	max_round = 100
	// max_x     = 11
	// max_y     = 7
	max_x = 101
	max_y = 103
)

type Robot struct {
	px int
	py int
	vx int
	vy int
}

func parseRobot(line string) Robot {
	all_groups := r.FindStringSubmatch(line)
	px, _ := strconv.Atoi(all_groups[1])
	py, _ := strconv.Atoi(all_groups[2])
	vx, _ := strconv.Atoi(all_groups[3])
	vy, _ := strconv.Atoi(all_groups[4])
	return Robot{px, py, vx, vy}
}

// It's ok Copilot made this
func stddev(input []int) float64 {
	sum := 0
	for _, i := range input {
		sum += i
	}
	mean := float64(sum) / float64(len(input))
	sum = 0
	for _, i := range input {
		sum += (i - int(mean)) * (i - int(mean))
	}
	return math.Sqrt(float64(sum) / float64(len(input)))
}

func main() {
	file_name := "input.txt"
	file_name = "real_input.txt"
	input := read_input(file_name)

	mid_x := max_x / 2
	mid_y := max_y / 2
	// fmt.Println(mid_x, mid_y)

	quads := []int{0, 0, 0, 0}
	robots := []Robot{}

	for _, line := range input {
		robot := parseRobot(line)
		// for part 2
		robots = append(robots, robot)

		final_x := ((robot.px+robot.vx*max_round)%max_x + max_x) % max_x
		final_y := ((robot.py+robot.vy*max_round)%max_y + max_y) % max_y

		if final_x == mid_x || final_y == mid_y {
			continue
		}

		// fmt.Println(final_x, final_y, final_x/mid_x, final_y/mid_y)
		quads[final_x/(mid_x+1)+2*(final_y/(mid_y+1))]++
	}

	fmt.Println(quads[0]*quads[1]*quads[2]*quads[3], quads)

	for i := 1; ; i++ {
		current_map := []Robot{}
		x_arr, y_arr := []int{}, []int{}

		for _, robot := range robots {
			next_px := (robot.px + robot.vx + max_x) % max_x
			next_py := (robot.py + robot.vy + max_y) % max_y
			next_robot := Robot{next_px, next_py, robot.vx, robot.vy}
			current_map = append(current_map, next_robot)
			x_arr = append(x_arr, next_px)
			y_arr = append(y_arr, next_py)
		}

		// You can see that most of the time,
		// the standard deviation of the robots are around 28~30.
		// If the robots are "aligned" to make the easter egg,
		// the standard deviation must be less than 25.
		// We can assume that's the answer.
		// ref: https://www.reddit.com/r/adventofcode/comments/1hdvhvu/comment/m1zgdsh/
		std_x := stddev(x_arr)
		std_y := stddev(y_arr)
		// fmt.Println(i, std_x, std_y)

		if std_x < 25 && std_y < 25 {
			// Extra
			final_map := make([][]string, max_y)
			for i := 0; i < max_y; i++ {
				final_map[i] = make([]string, max_x)
				for j := 0; j < max_x; j++ {
					final_map[i][j] = "."
				}
			}
			for _, robot := range current_map {
				final_map[robot.py][robot.px] = "#"
			}
			for i := 0; i < max_y; i++ {
				fmt.Println(strings.Join(final_map[i], ""))
			}
			// END Extra
			fmt.Println("FOUND", i, std_x, std_y)
			// FOUND i 19.231900582105762 20.314034557418672
			break
		}

		robots = current_map
	}
	fmt.Println("DONE")
}
