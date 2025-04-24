package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := range nums {

		if 0 < i && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}

		}
	}
	return ans
}

func compare(result, expected [][]int) bool {
	if len(result) != len(expected) {
		return false
	}

	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) > 0 && len(result[j]) > 0 {
			return result[i][0] < result[j][0]
		}
		return false
	})

	for i := range result {
		if len(result[i]) != len(expected[i]) {
			return false
		}
		for j := range result[i] {
			if result[i][j] != expected[i][j] {
				return false
			}
		}
	}

	return true
}
func main() {
	for i := range T.cases {

		ans := threeSum(T.cases[i])

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
