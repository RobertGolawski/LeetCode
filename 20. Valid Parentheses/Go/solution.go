package main

import (
	"fmt"
	"unicode/utf8"
)

func isValid(s string) bool {
	stack := []rune{}
	if utf8.RuneCountInString(s)%2 != 0 {
		return false
	}

	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}

		switch c {
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}

func main() {
	for i := range T.cases {

		str := T.cases[i]
		exp := T.answers[i]
		ans := isValid(str)

		passed := (ans == exp)

		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: %v \n", T.cases[i])
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Print("Congratulations")
}
