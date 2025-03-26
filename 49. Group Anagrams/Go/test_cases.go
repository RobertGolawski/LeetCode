package main

type tests struct {
	cases   [][]string
	answers []map[string]bool
}

var T = tests{
	cases: [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
		{"", ""},
		{"a", ""},
		{"a", "b"},
		{"ab", "ba"},
		{"ab", "ba", "cd", "dc", "e"},
		{"abc", "bac", "cba", "def", "fed"},
		{"listen", "silent"},
		{"rat", "tar", "art"},
	},
	answers: []map[string]bool{
		{
			"bat":         true,
			"nat,tan":     true,
			"ate,eat,tea": true,
		},
		{
			"": true,
		},
		{
			"a": true,
		},
		{
			",": true,
		},
		{
			"a": true,
			"":  true,
		},
		{
			"a": true,
			"b": true,
		},
		{
			"ab,ba": true,
		},
		{
			"ab,ba": true,
			"cd,dc": true,
			"e":     true,
		},
		{
			"abc,bac,cba": true,
			"def,fed":     true,
		},
		{
			"listen,silent": true,
		},
		{
			"art,rat,tar": true,
		},
	},
}
