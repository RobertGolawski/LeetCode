package main

type tests struct {
	cases   []*ListNode
	answers []*ListNode
}

var T = tests{
	cases: []*ListNode{
		// Test Case 1: An empty list.
		nil,
		// Test Case 2: A list with a single node.
		{Val: 1, Next: nil},
		// Test Case 3: A list with two nodes.
		{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
		// Test Case 4: A longer list of nodes.
		{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5, Next: nil},
					},
				},
			},
		},
		// Test Case 5: A list with negative numbers.
		{
			Val: -10,
			Next: &ListNode{
				Val:  25,
				Next: &ListNode{Val: -3, Next: nil},
			},
		},
		// Test Case 6: A list with duplicate values.
		{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
			},
		},
		// Test Case 7: A list containing a zero.
		{
			Val: 0,
			Next: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 10, Next: nil},
			},
		},
		// Test Case 8: A list with alternating positive and negative numbers.
		{
			Val: 1,
			Next: &ListNode{
				Val: -2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: -4, Next: nil},
				},
			},
		},
		// Test Case 9: A list where all elements are the same.
		{
			Val: 7,
			Next: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 7, Next: nil},
			},
		},
		// Test Case 10: A longer list.
		{
			Val: 10,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val:  6,
							Next: &ListNode{Val: 5, Next: nil},
						},
					},
				},
			},
		},
	},
	answers: []*ListNode{
		// Expected Answer 1: An empty list remains empty.
		nil,
		// Expected Answer 2: A single node list remains the same.
		{Val: 1, Next: nil},
		// Expected Answer 3: The two nodes should be reversed.
		{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
		// Expected Answer 4: The longer list should be fully reversed.
		{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 1, Next: nil},
					},
				},
			},
		},
		// Expected Answer 5: The list with negative numbers reversed.
		{
			Val: -3,
			Next: &ListNode{
				Val:  25,
				Next: &ListNode{Val: -10, Next: nil},
			},
		},
		// Expected Answer 6: The list with duplicates reversed.
		{
			Val: 3,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
			},
		},
		// Expected Answer 7: The list with a zero reversed.
		{
			Val: 10,
			Next: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 0, Next: nil},
			},
		},
		// Expected Answer 8: The list with alternating signs reversed.
		{
			Val: -4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  -2,
					Next: &ListNode{Val: 1, Next: nil},
				},
			},
		},
		// Expected Answer 9: The list with identical elements reversed.
		{
			Val: 7,
			Next: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 7, Next: nil},
			},
		},
		// Expected Answer 10: The longer list reversed.
		{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{Val: 10, Next: nil},
						},
					},
				},
			},
		},
	},
}

