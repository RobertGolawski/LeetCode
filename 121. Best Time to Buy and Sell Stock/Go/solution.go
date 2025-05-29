package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	res := 0
	l := 0
	r := 1

	for r < len(prices) {
		res = max(res, prices[r]-prices[l])
		if prices[r] < prices[l] {
			l = r
		}
		r++
	}
	return res
}

func main() {
	for i := range T.cases {

		ans := maxProfit(T.cases[i])

		passed := ans == T.answers[i]

		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: %v \n", T.cases)
			fmt.Printf("  Expected: %v\n", T.answers)
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Print("Congratulations")
}
