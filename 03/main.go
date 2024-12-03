package main

import (
	"bufio"
	"log"
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

func main2() {
	file_name := "real_input.txt"
	input := read_input(file_name)
	all_input := strings.Join(input, "")
	r := regexp.MustCompile(`do.*(mul\((\d+),(\d+)\).*)`)

	groups := [][]string{}
	for _, line := range input {
		pair := r.FindAllStringSubmatch(line, -1)
		for _, p := range pair {
			groups = append(groups, p)
		}
	}
	// fmt.Println(groups)
	result := 0
	for _, pair := range groups {
		a, b := pair[1], pair[2]
		a1, _ := strconv.Atoi(a)
		b1, _ := strconv.Atoi(b)
		result += a1 * b1
	}
	// fmt.Println(result)
}

func main() {
	file_name := "real_input.txt"
	input := read_input(file_name)
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	groups := [][]string{}
	for _, line := range input {
		pair := r.FindAllStringSubmatch(line, -1)
		for _, p := range pair {
			groups = append(groups, p)
		}
	}
	// fmt.Println(groups)
	result := 0
	for _, pair := range groups {
		a, b := pair[1], pair[2]
		a1, _ := strconv.Atoi(a)
		b1, _ := strconv.Atoi(b)
		result += a1 * b1
	}
	// fmt.Println(result)
}
