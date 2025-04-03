package main

type tests struct {
	cases   [][]int
	answers [][]int
}

var T = tests{
	cases: [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
		{1, 2, 3, 4, 5},
		{1},
		{0, 1},
		{2, 3},
		{-1, -2, -3},
	},
	answers: [][]int{
		{24, 12, 8, 6},
		{0, 0, 9, 0, 0},
		{120, 60, 40, 30, 24},
		{1},
		{1, 0},
		{3, 2},
		{6, 3, 2},
	},
}
