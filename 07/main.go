package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)

	result := 0
	for _, line := range input {
		target_str := line[:strings.Index(line, ":")]
		members_str := strings.Split(line[strings.Index(line, ":")+2:], " ")

		target, _ := strconv.Atoi(target_str)
		possible_target := []int{}

		for _, member_str := range members_str {
			member, _ := strconv.Atoi(member_str)
			if len(possible_target) == 0 {
				possible_target = append(possible_target, member)
				continue
			}
			next_target := []int{}
			for _, pt := range possible_target {
				// plus
				next_target = append(next_target, pt+member)
				// multiply
				next_target = append(next_target, pt*member)
				// concat
				pt_str := strconv.Itoa(pt) + strconv.Itoa(member)
				pt_int, _ := strconv.Atoi(pt_str)
				next_target = append(next_target, pt_int)
			}
			possible_target = next_target
		}

		if slices.Contains(possible_target, target) {
			result += target
			fmt.Println(target)
		}
	}

	fmt.Println(result)
}
