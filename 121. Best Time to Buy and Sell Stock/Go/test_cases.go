package main

type tests struct {
	cases   [][]int
	answers []int
}

var T = tests{
	cases: [][]int{
		{7, 1, 5, 3, 6, 4},       // Original test case
		{1, 2, 3, 4, 5},          // Prices strictly increasing
		{7, 6, 4, 3, 1},          // Prices strictly decreasing (no profit)
		{2, 4, 1},                // Buy at 2, sell at 4
		{3, 2, 6, 5, 0, 3},       // Buy at 2, sell at 6
		{1, 1, 1, 1, 1},          // Prices stay the same (no profit)
		{5},                      // Single day (no transaction possible)
		{},                       // Empty array (no transaction possible)
		{2, 1, 2, 1, 0, 1, 2},    // Buy at 0, sell at 2
		{3, 3, 5, 0, 0, 3, 1, 4}, // Buy at 0, sell at 4
		{10, 1, 9, 0, 8},         // Buy at 1, sell at 9 OR buy at 0, sell at 8
		{6, 5, 4, 3, 2, 1, 0, 1}, // Buy at 0, sell at 1
		{1, 2, 0, 3, 2, 4},       // Buy at 0, sell at 4
	},
	answers: []int{
		5, // Original test case: Buy at 1, sell at 6
		4, // {1, 2, 3, 4, 5}: Buy at 1, sell at 5
		0, // {7, 6, 4, 3, 1}: No profit
		2, // {2, 4, 1}: Buy at 2, sell at 4
		4, // {3, 2, 6, 5, 0, 3}: Buy at 2, sell at 6
		0, // {1, 1, 1, 1, 1}: No profit
		0, // {5}: No transaction
		0, // {}: No transaction
		2, // {2, 1, 2, 1, 0, 1, 2}: Buy at 0, sell at 2
		4, // {3, 3, 5, 0, 0, 3, 1, 4}: Buy at 0, sell at 4
		8, // {10, 1, 9, 0, 8}: Buy at 0, sell at 8 (max profit is 8, not 1 to 9 which is 8)
		1, // {6, 5, 4, 3, 2, 1, 0, 1}: Buy at 0, sell at 1
		4, // {1, 2, 0, 3, 2, 4}: Buy at 0, sell at 4
	},
}
