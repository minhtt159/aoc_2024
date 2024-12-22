package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var ex_input string

//go:embed input2.txt
var ex_input2 string

//go:embed real_input.txt
var real_input string

var (
	modulo = 16777216
	round  = 2000
)

func calcNext(num int) int {
	result := num
	next_a := result * 64
	result = (result ^ next_a) % modulo
	next_b := result / 32
	result = (result ^ next_b) % modulo
	next_c := result * 2048
	result = (result ^ next_c) % modulo
	return result
}

type Pair struct {
	line int
	num  int
}

func getHash(hash []int, i int) string {
	ret := []byte{}
	ret = append(ret, byte(hash[i-3]-hash[i-4]+'k'))
	ret = append(ret, byte(hash[i-2]-hash[i-3]+'k'))
	ret = append(ret, byte(hash[i-1]-hash[i-2]+'k'))
	ret = append(ret, byte(hash[i]-hash[i-1]+'k'))

	return string(ret)
}

func main() {
	input := strings.Split(real_input, "\n")

	sum := 0
	rep := make(map[string][]Pair)
	for j, line := range input {
		if len(line) == 0 {
			continue
		}
		num, _ := strconv.Atoi(line)
		hash := []int{}
		visited := make(map[string]bool)
		for i := 0; i < round; i++ {
			num = calcNext(num)
			h := num % 10
			hash = append(hash, h)
			if i > 4 {
				stage := getHash(hash, i)
				if _, ok := visited[stage]; !ok {
					// fmt.Println("Remembering", stage, h)
					rep[stage] = append(rep[stage], Pair{j, h})
				} else {
					// fmt.Println("Already stored", stage)
					continue
					// for _, v := range rep[stage] {
					// 	if v.line == j {
					// 		if h > v.num {
					// 			v.num = h
					// 		}
					// 	}
					// }
				}
				// fmt.Println("Caching", stage)
				visited[stage] = true
			}
		}
		sum += num
		// println(num)
	}
	fmt.Println(sum)

	fmt.Println(len(rep))
	res := 0
	for _, v := range rep {
		// fmt.Println(k)
		temp := 0
		for _, p := range v {
			// fmt.Println(k, p.num)
			temp += p.num
		}
		if temp > res {
			res = temp
		}
	}
	fmt.Println(res)
}
