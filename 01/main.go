package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// Read input from file
// There are 2 columns in each line
// Append them to 2 separate arrays
// Return the arrays
func read_input(file_name string) ([]int, []int) {
	a, b := []int{}, []int{}
	r := regexp.MustCompile(`(\d+)   (\d+)`)
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a1, b1 := r.FindStringSubmatch(line)[1], r.FindStringSubmatch(line)[2]
		fmt.Println(a1, b1)
		i1, err := strconv.Atoi(a1)
		if err != nil {
			log.Fatal(err)
		}
		i2, err := strconv.Atoi(b1)
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, i1)
		b = append(b, i2)
	}
	return a, b
}

func main() {
	file_name := "real_input.txt"
	a, b := read_input(file_name)
	sort.Ints(a)
	sort.Ints(b)
	// fmt.Println(a, b)

	result := 0
	for i := 0; i < len(a); i++ {
		diff := a[i] - b[i]
		fmt.Println(diff)
		if diff < 0 {
			result -= diff
		} else {
			result += diff
		}
	}
	fmt.Println(result)
}
