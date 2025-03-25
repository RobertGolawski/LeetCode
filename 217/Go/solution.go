package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
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
		ans := containsDuplicate(T.cases[i])
		if T.answers[i] != ans {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: %v\n", T.cases[i])
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Printf("Congratulations!")
}
