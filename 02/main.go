package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tribool "gopkg.in/grignaak/tribool.v1"
)

func read_input(file_name string) [][]string {
	result := [][]string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		result = append(result, levels)
	}
	return result
}

func arr_str_to_arr_int(a []string) []int {
	result := []int{}
	for i := 0; i < len(a); i++ {
		num, err := strconv.Atoi(a[i])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}
	return result
}

func eval(levels []string) []int {
	rep := []int{}
	l_int := arr_str_to_arr_int(levels)
	for i := 0; i < len(l_int)-1; i++ {
		rep = append(rep, l_int[i]-l_int[i+1])
	}
	return rep
}

func scan_report_01(levels []string) int {
	rep := eval(levels)
	// NOTE: this object is really easy to code, but I'm lazy
	all_neg := tribool.Maybe
	for i := 0; i < len(rep); i++ {
		if rep[i] == 0 || rep[i] < -3 || rep[i] > 3 {
			return 0
		} else if all_neg == tribool.Maybe {
			all_neg = tribool.FromBool(rep[i] < 0)
		} else if all_neg == tribool.True && rep[i] > 0 {
			return 0
		} else if all_neg == tribool.False && rep[i] < 0 {
			return 0
		}
		// NOTE: 2. I can obtimize this, but I'm also lazy
	}
	return 1
}

func scan_report_02(levels []string) int {
	if scan_report_01(levels) == 1 {
		return 1
	}
	attempt_result := 0
	for i := 0; i < len(levels); i++ {
		// Try to create a new slice without element i
		new_arr := append(append([]string{}, levels[:i]...), levels[i+1:]...)
		attempt_result |= scan_report_01(new_arr)
	}
	return attempt_result
}

func main() {
	file_name := "real_input.txt"
	a := read_input(file_name)
	result := 0

	for i := 0; i < len(a); i++ {
		// fmt.Println(a[i])
		// result += scan_report_01(a[i])
		result += scan_report_02(a[i])
	}

	fmt.Println(result)
}
