package main

import (
	_ "embed"
	"strconv"
)

// 029A <- 29 * 68
// 980A <- 980 * 60
// 179A <- 179 * 68
// 456A <- 456 * 64
// 379A <- 379 * 64

//go:embed input.txt
var ex_input string

//go:embed real_input.txt
var real_input string

const (
	U = iota
	L
	D
	R
	N = -1
	A = 10
)

type Case struct {
	input string
	round int
}

var (
	// num pad
	num_map = [][3]int{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
		{N, 0, A},
	}
	// direction pad
	dir_map = [][3]int{
		{N, U, A},
		{L, D, R},
	}
	dir_x = []int{-1, 0, 1, 0}
	dir_y = []int{0, -1, 0, 1}
	// route cache
	case1 = Case{
		ex_input,
		2,
	}
	case2 = Case{
		ex_input,
		25,
	}
	case3 = Case{
		real_input,
		2,
	}
	case4 = Case{
		real_input,
		25,
	}
)

// Given a rune, find the position in the map
func findPosInMap(input [][3]int, r rune) Pos {
	var n1 int
	switch r {
	case 'A':
		n1 = A
	case 'N':
		n1 = N
	case '^':
		n1 = U
	case '<':
		n1 = L
	case 'v':
		n1 = D
	case '>':
		n1 = R
	default:
		n1, _ = strconv.Atoi(string(r))
	}
	for i, row := range input {
		for j, n := range row {
			if n == n1 {
				return Pos{i, j}
			}
		}
	}
	return Pos{-1, -1}
}
