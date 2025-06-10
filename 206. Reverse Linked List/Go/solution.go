package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func areListsEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	// If both are nil, the lists are equal in length and content.
	return l1 == nil && l2 == nil
}
func main() {
	for i := range T.cases {
		inputCopy := T.cases[i]
		got := reverseList(inputCopy)
		expected := T.answers[i]

		passed := areListsEqual(got, expected)
		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: %v \n", T.cases)
			fmt.Printf("  Expected: %v\n", T.answers)
			fmt.Printf("  Got: %v\n", got)
			return
		}
	}
	fmt.Print("Congratulations")
}
