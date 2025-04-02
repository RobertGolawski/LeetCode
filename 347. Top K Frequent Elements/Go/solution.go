package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, n := range nums {
		freqMap[n]++
	}

	freqList := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		freqList[freq] = append(freqList[freq], num)
	}

	result := []int{}
	for i := len(freqList) - 1; i > 0 && len(result) < k; i-- {
		result = append(result, freqList[i]...)
	}

	return result[:k]
}

func compare(ans, ret []int) bool {
	if len(ans) != len(ret) {
		return false
	}
	sort.Ints(ans)
	sort.Ints(ret)

	for i := range ans {
		if ans[i] != ret[i] {
			return false
		}
	}
	return true
}

func main() {
	for i := range T.cases {
		nums := T.cases[i][:len(T.cases[i])-1]
		k := T.cases[i][len(T.cases[i])-1]

		ans := topKFrequent(nums, k)

		passed := compare(T.answers[i], ans)

		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Nums: %v K: %v\n", nums, k)
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Print("Congratulations")
}
