package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
)

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

var small_input = `
x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02
`

func parseInput(input string) (map[string]int, [][]string) {
	r1 := regexp.MustCompile(`(.*): (\d)`)
	r2 := regexp.MustCompile(`(.*) (AND|OR|XOR) (.*) -> (.*)`)
	listConns := r1.FindAllStringSubmatch(input, -1)
	listGates := r2.FindAllStringSubmatch(input, -1)
	// listGates := [][]string{}
	// for _, line := range strings.Split(input, "\n") {
	// 	if G := r2.FindAllStringSubmatch(line, -1); G != nil {
	// 		listGates = append(listGates, G[0])
	// 	}
	// }

	connMap := make(map[string]int)
	for _, conn := range listConns {
		if conn[2] == "1" {
			connMap[conn[1]] = 1
		} else {
			connMap[conn[1]] = 0
		}
	}

	return connMap, listGates
}

func main() {
	// input := ex_input
	// input := small_input
	input := real_input
	mapConn, listGates := parseInput(input)
	// for _, gate := range listGates {
	// 	fmt.Println(gate)
	// }

	i := 0
	for len(listGates) > 0 {
		i++
		if i >= len(listGates) {
			i = 0
		}
		thisGate := listGates[i]
		if _, ok := mapConn[thisGate[1]]; !ok {
			continue
		} else if _, ok := mapConn[thisGate[3]]; !ok {
			continue
		}
		switch thisGate[2] {
		case "AND":
			mapConn[thisGate[4]] = mapConn[thisGate[1]] & mapConn[thisGate[3]]
		case "OR":
			mapConn[thisGate[4]] = mapConn[thisGate[1]] | mapConn[thisGate[3]]
		case "XOR":
			mapConn[thisGate[4]] = mapConn[thisGate[1]] ^ mapConn[thisGate[3]]
		}
		listGates = append(listGates[:i], listGates[i+1:]...)
	}

	sortedConnKey := []string{}
	for k := range mapConn {
		// Get all gates starts with z
		sortedConnKey = append(sortedConnKey, k)
	}
	sort.Strings(sortedConnKey)
	slices.Reverse(sortedConnKey)

	ret := 0
	for _, k := range sortedConnKey {
		if k[0] != 'z' {
			continue
		}
		fmt.Println(k, mapConn[k])
		ret += mapConn[k]
		ret <<= 1
	}
	fmt.Println(strconv.FormatInt(int64(ret), 2), ret>>1)
}
