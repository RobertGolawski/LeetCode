package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := make([]int, 26)
	b := make([]int, 26)

	for i := range s {
		a[rune(s[i])-rune('a')] += 1
		b[rune(t[i])-rune('a')] += 1
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	for i := range T.cases {
		if T.answers[i] != isAnagram(T.cases[i][0], T.cases[i][1]) {
			fmt.Printf("Incorrect answer in test case %d", i)
			return
		}
	}
	fmt.Printf("Congratulations!")
}
