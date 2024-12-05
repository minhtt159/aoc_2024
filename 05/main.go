package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func CheckLine(tokens []int, rules map[int][]int) (bool, int) {
	// Loop through each rule
	// fmt.Println(tokens)
	for k, rule := range rules {
		// fmt.Println(k)
		if !slices.Contains(tokens, k) {
			// Rule does not apply for this token
			continue
		}
		for _, v := range rule {
			// a is the index of rule key in tokens
			a := slices.Index(tokens, k)
			// b is the index of rule value in tokens
			b := slices.Index(tokens, v)
			// fmt.Println(k, v, a, b)
			if a > b && a != -1 && b != -1 {
				return false, 0
			}
		}
	}
	return true, tokens[len(tokens)/2]
}

func FixLine(tokens []int, rules map[int][]int) int {
	// fmt.Println(tokens)
	for {
		for k, rule := range rules {
			if !slices.Contains(tokens, k) {
				continue
			}
			for _, v := range rule {
				// a is the index of rule key in tokens
				a := slices.Index(tokens, k)
				// b is the index of rule value in tokens
				b := slices.Index(tokens, v)
				// fmt.Println(k, v, a, b)
				if a > b && a != -1 && b != -1 {
					tokens[a], tokens[b] = tokens[b], tokens[a]
				}
			}
		}
		if isComply, mid := CheckLine(tokens, rules); isComply {
			return mid
		}
	}
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	r_order := regexp.MustCompile(`(\d+)|(\d+)`)

	var line_number int
	rules := make(map[int][]int)

	for line_number = 0; line_number < len(input); line_number++ {
		ret := r_order.FindAllString(input[line_number], -1)
		if len(ret) == 0 {
			break
		}
		// fmt.Println(ret)

		a, _ := strconv.Atoi(ret[0])
		b, _ := strconv.Atoi(ret[1])

		// Add b to rules[a]
		rules[a] = append(rules[a], b)
	}
	// fmt.Println(rules)

	result := 0
	result_incorrect := 0
	for ; line_number < len(input); line_number++ {
		intTokens := []int{}
		for _, str := range strings.Split(input[line_number], ",") {
			num, _ := strconv.Atoi(str)
			intTokens = append(intTokens, num)
		}
		if isComply, mid := CheckLine(intTokens, rules); isComply {
			result += mid
		} else {
			// Try to order them
			result_incorrect += FixLine(intTokens, rules)
		}
	}
	fmt.Println(result)
	fmt.Println(result_incorrect)
}
