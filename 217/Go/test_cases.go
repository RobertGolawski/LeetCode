package main

type tests struct {
	cases   [][]int
	answers []bool
}

var T = tests{
	cases: [][]int{
		{1, 2, 3, 1},                    // Basic duplicate
		{1, 2, 3, 4},                    // No duplicate
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},  // Multiple duplicates
		{},                              // Empty array
		{1},                             // Single element
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // No duplicates, longer array
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},  // Duplicate at the end
		{1, 2, 3, 4, 5, 6, 7, 8, 1, 9},  // Duplicate in the middle
		{-1, -2, -3, -1},                // Negative numbers, duplicate
		{-1, -2, -3, -4},                // Negative numbers, no duplicate
		{0, 0, 0, 0},                    // All zeros
		{1, 2, 1, 2, 1, 2},              // Alternating duplicates
		{1, 2, 3, 1, 2, 3},              // Duplicate sequences
		{100000, 100000, 100000},        // Large numbers, duplicate
		{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},  // Sorted with adjacent duplicates
		{5, 4, 3, 2, 1, 5},              // Unsorted with duplicate
		{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},  // multiple sequence duplicates
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},  // Duplicate at end longer array
	},
	answers: []bool{
		true,
		false,
		true,
		false,
		false,
		false,
		true,
		true,
		true,
		false,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
	},
}
