package main

type tests struct {
	cases   [][]string
	answers []bool
}

var T = tests{
	cases:   [][]string{{"anagram", "nagaram"}, {"rat", "car"}},
	answers: []bool{true, false},
}
