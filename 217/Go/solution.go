package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, i := range nums {
		if _, ok := m[i]; ok {
			return true
		}
		m[i] = true
	}
	return false
}

func main() {
	for i := range T.cases {
		if T.answers[i] != containsDuplicate(T.cases[i]) {
			fmt.Printf("Incorrect answer in test case %d", i)
			return
		}
	}
	fmt.Printf("Congratulations!")
}
