package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

var corrupted_input = `Register A: 117440
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

var test_input = `Register A: 37221270076916
Register B: 0
Register C: 0

Program: 2,4,1,2,7,5,4,5,1,3,5,5,0,3,3,0`

type Register struct {
	output []int64
	a      int64
	b      int64
	c      int64
}

type Program struct {
	opcode  int
	operand int
}

const (
	adv = iota // rA = rA / 2 ** x
	bxl        // rB = rB ^ x
	bst        // rB = x % 8
	jnz        // if rA == 0, do nothing | else jump to literal operand
	bxc        // rB = rB ^ rC
	out        // print(x)
	bdv        // rB = rA / 2 ** x
	cdv        // rC = rA / 2 ** x
)

var max_i = 8

func parseInput(input string) (Register, []Program, string) {
	all_lines := strings.Split(input, "\n")
	rA, _ := strconv.Atoi(all_lines[0][12:])
	rB, _ := strconv.Atoi(all_lines[1][12:])
	rC, _ := strconv.Atoi(all_lines[2][12:])

	p := strings.Split(all_lines[4][9:], ",")
	all_p := []Program{}
	for i := 0; i < len(p); i += 2 {
		opcode, _ := strconv.Atoi(p[i])
		operand, _ := strconv.Atoi(p[i+1])
		all_p = append(all_p, Program{opcode, operand})
	}
	return Register{a: int64(rA), b: int64(rB), c: int64(rC), output: []int64{}}, all_p, all_lines[4][9:]
}

func getOperand(r *Register, combo int) int64 {
	switch combo {
	case 0, 1, 2, 3:
		return int64(combo)
	case 4:
		return r.a
	case 5:
		return r.b
	case 6:
		return r.c
	default:
		log.Fatal("Invalid Operand")
		return -1
	}
}

func Fadv(r *Register, p Program) (bool, int) {
	x := getOperand(r, p.operand)
	deno := int64(math.Pow(2, float64(x)))
	r.a = int64(r.a / deno)
	return true, -1
}

func Fbdv(r *Register, p Program) (bool, int) {
	x := getOperand(r, p.operand)
	deno := int64(math.Pow(2, float64(x)))
	r.b = int64(r.a / deno)
	return true, -1
}

func Fcdv(r *Register, p Program) (bool, int) {
	x := getOperand(r, p.operand)
	deno := int64(math.Pow(2, float64(x)))
	r.c = int64(r.a / deno)
	return true, -1
}

func Fout(r *Register, p Program) (bool, int) {
	x := getOperand(r, p.operand) % 8 // r.hidden = r.b % 8
	r.output = append(r.output, x)
	// fmt.Println("DEBUG:", x)
	return true, -1
}

func Fbxl(r *Register, p Program) (bool, int) {
	x := getOperand(r, p.operand)
	r.b = r.b ^ x
	return true, -1
}

func Fbxc(r *Register, p Program) (bool, int) {
	r.b = r.b ^ r.c
	return true, -1
}

func Fbst(r *Register, p Program) (bool, int) {
	x := getOperand(r, p.operand)
	r.b = x % 8
	return true, -1
}

func Fjnz(r *Register, p Program) (bool, int) {
	if r.a == 0 {
		return true, -1
	} else {
		return false, p.operand / 2
	}
}

func run(register Register, programs []Program) []int64 {
	list_opcode := []func(*Register, Program) (bool, int){Fadv, Fbxl, Fbst, Fjnz, Fbxc, Fout, Fbdv, Fcdv}

	i := 0
	for {
		p := programs[i]
		f := list_opcode[p.opcode]
		// fmt.Println(register, p)
		if ok, val := f(&register, p); !ok {
			i = val
		} else {
			i = i + 1
		}
		if i >= len(programs) {
			break
		}
	}
	return register.output
}

func compare(x []int64, y []int64) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	file := ex_input
	file = real_input
	// file = corrupted_input
	file = test_input
	register, programs, program_str := parseInput(file)
	// fmt.Println(register, programs)

	// part 1
	fmt.Println("p1", run(register, programs))

	// Analyse program
	// [{2 4} {1 2} {7 5} {4 5} {1 3} {5 5} {0 3} {3 0}
	// 2,4 | rB = rA % 8  <-- rB is the last 3 bit of rA
	// 1,2 | rB = rB ^ 2
	// 7,5 | rC = rA / 2 ** rB
	// 4,5 | rB = rB ^ rC
	// 1,3 | rB = rB ^ 3
	// 5,5 | out(rB % 8)
	// 0,3 | rA = rA / 8  <- basically throw away the last 3 bit
	// part 2
	// 3,0 | goto 1
	// On each output, it only cares about the last 3 bit of rA

	target := []int64{}
	for _, t := range strings.Split(program_str, ",") {
		t_int, _ := strconv.Atoi(t)
		target = append(target, int64(t_int))
	}
	// fmt.Println("Target", target)

	potentials := []int64{int64(0)}
	for i := len(target) - 1; i >= 0; i-- {
		new_po := []int64{}

		for _, a := range potentials {
			a <<= 3
			// Brute last 3 digit of rA
			for b := int64(0); b < 8; b++ {
				potential := a + b
				this_register := Register{a: potential, b: 0, c: 0, output: []int64{}}
				r := run(this_register, programs)
				// fmt.Println("--", potential, strconv.FormatInt(potential, 8), r, target[i:])

				if compare(r, target[i:]) {
					// a = potential
					// break
					new_po = append(new_po, potential)
				}
			}
		}

		potentials = new_po
	}
	fmt.Println("p2", potentials)
}
