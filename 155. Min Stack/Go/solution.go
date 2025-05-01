package main

import (
	"fmt"
)

type MinStack struct {
	min []int
	arr []int
}

func Constructor() MinStack {
	return MinStack{
		min: []int{},
		arr: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)

	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		return
	}
	this.arr = this.arr[:len(this.arr)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

func main() {
	for i, tc := range T { // Iterate through each test case
		var stack MinStack
		var result interface{} // To store results from Top/GetMin

		fmt.Printf("Running Test Case %d...\n", i)

		for j, command := range tc.commands { // Iterate through commands in the test case
			args := tc.arguments[j]
			expected := tc.expected[j]
			opFailed := false

			// Execute the command
			switch command {
			case "Constructor":
				stack = Constructor()
				result = nil // Constructor doesn't return a value to check here
			case "Push":
				if len(args) == 1 {
					stack.Push(args[0])
					result = nil // Push doesn't return a value to check here
				} else {
					fmt.Printf("  Error: Incorrect arguments for Push in test case %d, step %d\n", i, j)
					opFailed = true
				}
			case "Pop":
				stack.Pop()
				result = nil // Pop doesn't return a value to check here
			case "Top":
				result = stack.Top()
			case "GetMin":
				result = stack.GetMin()
			default:
				fmt.Printf("  Error: Unknown command '%s' in test case %d, step %d\n", command, i, j)
				opFailed = true
			}

			if opFailed {
				return // Stop processing this test case if an error occurred
			}

			// Check the result if an expected value is provided (not nil)
			if expected != nil {
				if result != expected {
					fmt.Printf("Test Case %d Failed at Step %d (%s):\n", i, j, command)
					fmt.Printf("  Arguments: %v\n", args)
					fmt.Printf("  Expected: %v\n", expected)
					fmt.Printf("  Got:      %v\n", result)
					// Optional: Print stack state for debugging
					// fmt.Printf("  Stack arr: %v\n", stack.arr)
					// fmt.Printf("  Stack min: %v\n", stack.min)
					return // Stop on first failure
				}
			}
			// Optional: Print step success for debugging
			// else {
			//  fmt.Printf("  Step %d (%s) OK\n", j, command)
			// }
		}
		fmt.Printf("Test Case %d Passed.\n", i)
	}
	fmt.Println("Congratulations! All test cases passed.")
}
