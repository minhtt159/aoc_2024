package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Read input from file
// There are 2 columns in each line
// Append them to 2 separate arrays
// Return the arrays
func read_input(file_name string) ([]string, []string) {
	a, b := []string{}, []string{}
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

		a = append(a, a1)
		b = append(b, b1)
	}
	return a, b
}

func main_01() {
	file_name := "real_input.txt"
	a, b := read_input(file_name)
	a_int, b_int := []int{}, []int{}
	for i := 0; i < len(a); i++ {
		i1, err := strconv.Atoi(a[i])
		if err != nil {
			log.Fatal(err)
		}
		a_int = append(a_int, i1)
		i2, err := strconv.Atoi(b[i])
		if err != nil {
			log.Fatal(err)
		}
		b_int = append(b_int, i2)
	}
	sort.Ints(a_int)
	sort.Ints(b_int)

	result := 0
	for i := 0; i < len(a_int); i++ {
		diff := a_int[i] - b_int[i]
		if diff < 0 {
			result -= diff
		} else {
			result += diff
		}
	}
	fmt.Println(result)
}

func main_02() {
	file_name := "real_input.txt"
	a, b := read_input(file_name)
	b_str := strings.Join(b, ", ")
	result := 0

	for i := 0; i < len(a); i++ {
		r := regexp.MustCompile(regexp.QuoteMeta(a[i]))
		cnt := len(r.FindAllStringIndex(b_str, -1))
		base, err := strconv.Atoi(a[i])
		if err != nil {
			log.Fatal(err)
		}
		result += base * cnt
	}

	fmt.Println(result)
}

func main() {
	main_01()
	main_02()
}
