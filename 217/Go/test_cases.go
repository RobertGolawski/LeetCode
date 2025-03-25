package main

type tests struct {
	cases   [][]int
	answers []bool
}

var T = tests{
	cases:   [][]int{{1, 2, 3, 1}, {1, 2, 3, 4}, {1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
	answers: []bool{true, false, true},
}
