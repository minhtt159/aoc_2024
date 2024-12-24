package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"unicode/utf8"
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

var small_input2 = `
x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00
`

type Pair struct {
	a  string
	b  string
	op string
}

func parseInput(input string) (map[string]int, map[string]Pair) {
	r1 := regexp.MustCompile(`(.*): (\d)`)
	r2 := regexp.MustCompile(`(.*) (AND|OR|XOR) (.*) -> (.*)`)
	listConns := r1.FindAllStringSubmatch(input, -1)
	listGates := r2.FindAllStringSubmatch(input, -1)

	connMap := make(map[string]int)
	for _, conn := range listConns {
		if conn[2] == "1" {
			connMap[conn[1]] = 1
		} else {
			connMap[conn[1]] = 0
		}
	}

	gateMap := make(map[string]Pair)
	for _, gate := range listGates {
		thisGate := Pair{gate[1], gate[3], gate[2]}
		gateMap[gate[4]] = thisGate
	}

	return connMap, gateMap
}

func getVal(gate string, connMap map[string]int, gateMap map[string]Pair) int {
	pair := gateMap[gate]
	a, b := pair.a, pair.b
	var val_a, val_b int

	if gate_b, ok := connMap[b]; !ok {
		val_b = getVal(b, connMap, gateMap)
	} else {
		val_b = gate_b
	}

	if gate_a, ok := connMap[a]; !ok {
		val_a = getVal(a, connMap, gateMap)
	} else {
		val_a = gate_a
	}

	switch pair.op {
	case "AND":
		return val_a & val_b
	case "OR":
		return val_a | val_b
	case "XOR":
		return val_a ^ val_b
	default:
		log.Fatalf("Unknown operator: %s", pair.op)
		return -1
	}
}

func convert(token byte, val map[string]int) (int, string) {
	arr_token := []string{}
	for k := range val {
		// ret = string(val[k]+48) + ret
		if k[0] == token {
			arr_token = append(arr_token, k)
		}
	}
	slices.Sort(arr_token)

	ret_str := ""
	for _, k := range arr_token {
		ret_str = string(val[k]+48) + ret_str
	}
	ret_int, _ := strconv.ParseInt(ret_str, 2, 64)
	return int(ret_int), ret_str
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func main() {
	// input := ex_input
	input := small_input2
	// input := real_input
	mapConn, listGates := parseInput(input)

	ret := make(map[string]int)
	for k := range listGates {
		ret[k] = getVal(k, mapConn, listGates)
	}

	int_z, str_z := convert('z', ret)
	fmt.Println("Part 1:", int_z, str_z)

	// part 2
	int_x, str_x := convert('x', mapConn)
	int_y, str_y := convert('y', mapConn)
	fmt.Println(int_x, str_x)
	fmt.Println(int_y, str_y)
	int_z2 := int_x + int_y
	str_z2 := Reverse(strconv.FormatInt(int64(int_z2), 2))
	fmt.Println(int_z2, str_z2)

	for i := range ret {
		index, _ := strconv.Atoi(i[1:])
		if index > len(str_z2) || index > len(str_z) {
			continue
		}
		if str_z2[index] != str_z[index] {
			fmt.Println(index, str_z2[index], str_z[index])
		}
	}
}

// NOTE: https://en.wikipedia.org/wiki/Adder_(electronics)#Ripple-carry_adder
func rca(i int, listGates map[string]Pair, swapped []string) (int, int) {
}
