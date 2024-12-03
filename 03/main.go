package main

import (
	"bufio"
	"fmt"
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

func main02() {
	file_name := "real_input.txt"
	input := read_input(file_name)
	input = append([]string{"do()"}, input...)
	all_input := strings.Join(input, "")
	new_input := strings.Split(all_input, "do()")
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	result := 0
	for _, line := range new_input {
		index := strings.Index(line, "don't")
		if index == -1 {
			index = len(line)
		}
		do := line[:index]
		// fmt.Println(do)
		nice := r.FindAllStringSubmatch(do, -1)
		for _, n := range nice {
			// fmt.Println(n)
			a, _ := strconv.Atoi(n[1])
			b, _ := strconv.Atoi(n[2])
			result += a * b
		}
	}
	fmt.Println(result)
}

func main01() {
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

func main() {
	main01()
	main02()
}
