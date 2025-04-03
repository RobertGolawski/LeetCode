package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	for i := range numbers {
		searchFor := target - numbers[i]
		left := i + 1
		right := len(numbers) - 1
		for left <= right {
			mid := left + (right-left)/2
			if numbers[mid] == searchFor {
				return []int{i + 1, mid + 1}
			}
			if searchFor < numbers[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}
	return []int{}
}

func compare(ans, ret []int) bool {
	if len(ans) != len(ret) {
		return false
	}
	if len(ret) != 2 {
		return false
	}

	for i := range ans {
		if ans[i] != ret[i] {
			return false
		}
	}
	return true
}

func main() {
	for i := range T.cases {
		numbers := T.cases[i][:len(T.cases[i])-1]
		target := T.cases[i][len(T.cases[i])-1]
		ans := twoSum(numbers, target)

		passed := compare(T.answers[i], ans)

		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Numbers: %v Target: %v \n", T.cases[i], target)
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Print("Congratulations")
}
