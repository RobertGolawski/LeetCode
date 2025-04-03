package main

import (
	"fmt"
)

func templateFunc() bool {
	return true
}

func main() {
	for i := range T.cases {

		ans := templateFunc()

		passed := ans

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
