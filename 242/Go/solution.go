package main

import (
	"fmt"
)

func isAnagram(s, t string) bool {
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
		ans := isAnagram(T.cases[i][0], T.cases[i][1])
		if T.answers[i] != ans {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: s=%v, t=%v\n", T.cases[i][0], T.cases[i][1])
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Printf("Congratulations!")
}
