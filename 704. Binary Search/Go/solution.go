package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		m := nums[mid]
		if m == target {
			return mid
		}

		if m < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	// Iterate through each test case defined in test_cases.go
	for i, tc := range T {
		// Call the function being tested
		result := search(tc.nums, tc.target)

		// Compare the actual result with the expected answer
		passed := (result == tc.answer)

		// If the result doesn't match the expected answer, print details and exit
		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input nums:   %v\n", tc.nums)
			fmt.Printf("  Input target: %d\n", tc.target)
			fmt.Printf("  Expected:     %d\n", tc.answer)
			fmt.Printf("  Got:          %d\n", result)
			return // Stop execution on the first failure
		}
	}

	// If all test cases passed, print a success message
	fmt.Println("Congratulations! All test cases passed.")
}
