package main

type tests struct {
	cases []struct {
		nums   []int
		target int
	}
	answers []map[int]bool // Now using map[int]bool for indices
}

var T = tests{
	cases: []struct {
		nums   []int
		target int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9},         // Basic test case
		{nums: []int{3, 2, 4}, target: 6},              // Another basic test case
		{nums: []int{3, 3}, target: 6},                 // Duplicate numbers that sum to target
		{nums: []int{0, 4, 3, 0}, target: 0},           // Zeros in the array
		{nums: []int{-1, -3, 6, 2}, target: 8},         // Negative numbers and positive target
		{nums: []int{-3, 4, 3, 90}, target: 0},         // Negative and positive numbers
		{nums: []int{5, 5, 5, 5}, target: 10},          // All same numbers, target is sum of two
		{nums: []int{1, 2, 3, 4, 5}, target: 1},        // No solution exists
		{nums: []int{1, 2, 3, 4, 5}, target: 10},       // Target larger than any possible sum of two
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, target: 13}, // Longer array
		{nums: []int{2, 5, 5, 11}, target: 10},         // Example with multiple 5's, ensure correct indices are returned
	},
	answers: []map[int]bool{
		{0: true, 1: true}, // indices for {2, 7}, target 9
		{1: true, 2: true}, // indices for {2, 4}, target 6
		{0: true, 1: true}, // indices for {3, 3}, target 6 (duplicate value is fine)
		{0: true, 3: true}, // indices for {0, 0}, target 0 (duplicate value is fine)
		{2: true, 3: true}, // indices for {6, 2}, target 8
		{0: true, 2: true}, // indices for {-3, 3}, target 0
		{0: true, 1: true}, // indices for {5, 5}, target 10
		{},                 // No solution for target 1
		{},                 // No solution for target 10
		{5: true, 6: true}, // indices for Target 13 in longer array {6,7}
		{1: true, 2: true}, // Indices for case where multiple identical integers result in a valid result
	},
}
