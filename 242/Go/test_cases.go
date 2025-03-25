package main

type tests struct {
	cases   [][]string
	answers []bool
}

var T = tests{
	cases: [][]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"", ""},
		{"a", "a"},
		{"a", "b"},
		{"ab", "ba"},
		{"ab", "ca"},
		{"abc", "cba"},
		{"abc", "abd"},
		{"listen", "silent"},
		{"triangle", "integral"},
		{"hello", "olleh"},
		{"abcabc", "bcabca"},
		{"abcabc", "aabbcc"},
		{"abcabc", "acbacb"},
		{"abcabc", "acccab"},
		{"abc", "abcd"},
		{"abcd", "abc"},
		{"aab", "aba"},
		{"aba", "baa"},
	},
	answers: []bool{
		true,
		false,
		true,
		true,
		false,
		true,
		false,
		true,
		false,
		true,
		true,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		true,
		true,
	},
}
