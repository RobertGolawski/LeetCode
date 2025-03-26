package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{i, v}
		}
		m[n] = i
	}
	return []int{}
}

func main() {
	for i := range T.cases {
		aSet := T.answers[i]
		res := twoSum(T.cases[i].nums, T.cases[i].target)
		var aOk bool
		var bOk bool
		if len(aSet) == 0 && len(res) == 0 {
			continue
		} else if len(res) == 2 {
			_, aOk = aSet[res[0]]
			_, bOk = aSet[res[1]]
		}

		if !aOk || !bOk {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: nums=%v, target=%d\n", T.cases[i].nums, T.cases[i].target)
			fmt.Printf("  Expected: %v\n", aSet)
			fmt.Printf("  Got: %v\n", res)
			return
		}
	}
	fmt.Print("Congratulations")
}
