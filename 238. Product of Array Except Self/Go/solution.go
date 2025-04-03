package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	ans := make([]int, len(nums)-1)

	p := 1
	ans = append(ans, nums[0])
	for i := 1; i < len(nums); i++ {
		ans[i] = p * nums[i-1]
		p = ans[i]
	}

	p2 := 1
	for i := len(nums) - 2; i > -1; i-- {
		p2 = p2 * nums[i+1]
		ans[i] = ans[i] * p2
	}
	ans[0] = p2

	return ans
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
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

		ans := productExceptSelf(T.cases[i])
		expected := T.answers[i]

		passed := compare(ans, expected)

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
