package main

type tests struct {
	cases   []string
	answers []bool
}

var T = tests{
	cases: []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
		"0P",
		"Was it a car or a cat I saw?",
		"Madam, I'm Adam!",
		"No 'x' in Nixon",
		"hello",
		"12321",
		"12345",
		"Able was I, ere I saw Elba!",
		"Kayak",
		"A",
		"",
	},
	answers: []bool{
		true,
		false,
		true,
		false,
		true,
		true,
		true,
		false,
		true,
		false,
		true,
		true,
		true,
		true,
	},
}
