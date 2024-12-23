package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

var network = make(map[string][]string)

type Group struct {
	comp [3]string
}

func parseInput(input string) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		conns := strings.Split(line, "-")
		if _, ok := network[conns[0]]; !ok {
			network[conns[0]] = []string{}
		}
		network[conns[0]] = append(network[conns[0]], conns[1])
		if _, ok := network[conns[1]]; !ok {
			network[conns[1]] = []string{}
		}
		network[conns[1]] = append(network[conns[1]], conns[0])
	}
}

func BronKerbosch(node string) {
}

func main() {
	// input := ex_input
	input := real_input
	parseInput(input)

	ret := make(map[Group]bool)

	for a := range network {
		// fmt.Println(comp, network[comp])
		if a[0] == 't' {
			for _, b := range network[a] {
				for _, c := range network[b] {
					if slices.Contains(network[a], c) {
						temp := []string{a, b, c}
						slices.Sort(temp)
						temp2 := [3]string{temp[0], temp[1], temp[2]}
						group := Group{temp2}
						if _, ok := ret[group]; !ok {
							ret[group] = true
						}
					}
				}
			}
		}
	}
	fmt.Println(len(ret))

	cliques := findCliques(network)
	r2 := []string{}
	for _, clique := range cliques {
		fmt.Println(clique)
		if len(r2) < len(clique) {
			r2 = clique
			slices.Sort(r2)
		}
	}
	fmt.Println(r2)
	fmt.Println(strings.Join(r2, ","))
}
