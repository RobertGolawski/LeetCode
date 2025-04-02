package main

type tests struct {
	cases   [][]int
	answers [][]int
}

var T = tests{
	cases: [][]int{
		{1, 1, 1, 2, 2, 3, 2},             // Test case 1
		{1, 2, 1},                         // Test case 2
		{1, 2, 3, 1},                      // Test case 3
		{1, 1, 2, 2, 3, 3, 3},             // Test case 4
		{4, 1, -1, 2, -1, 2, 3, 2},        // Test case 5
		{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4}, // Test case 6
		{1, 2, 2, 3, 3, 3, 2},             // Test case 7
		{1, 2, 3, 4, 5, 6, 3},             // Test case 8
		{1, 1, 1, 1, 2, 2, 2, 3, 3, 2},    // Test case 9
	},
	answers: [][]int{ // List of test results (array of arrays)
		{1, 2},       // Correct output for case 1
		{1},          // Correct output for case 2
		{1},          // Correct output for case 3. Changed from {1, 2, 3} to {1} as expected output
		{3, 1, 2},    // Correct output for case 4
		{-1, 2},      // Correct output for case 5
		{4, 1, 2, 3}, // Correct output for case 6
		{3, 2},       // Correct output for case 7
		{1, 2, 3},    // Correct output for case 8
		{1, 2},       // Correct output for case 9
	},
}
