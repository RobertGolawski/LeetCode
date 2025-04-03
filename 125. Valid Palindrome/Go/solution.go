package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	trimmed := strings.Map(func(r rune) rune {
		if isLetter(r) || isDigit(r) {
			return r
		}
		return -1
	}, s)
	trimmed = strings.ToLower(trimmed)

	runes := []rune(trimmed)
	left := 0
	right := len(runes) - 1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z' || ('A' <= r && r <= 'Z'))
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func main() {
	for i := range T.cases {

		ans := isPalindrome(T.cases[i])

		if ans != T.answers[i] {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: %v \n", T.cases[i])
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Print("Congratulations")
}
