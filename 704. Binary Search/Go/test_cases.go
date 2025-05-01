package main

type testCase struct {
	nums   []int // Input array (must be sorted for binary search)
	target int   // The value to search for
	answer int   // Expected index of the target, or -1 if not found
}

// T holds all the test cases
var T = []testCase{
	// Test Case 0: Target found in the middle
	{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 9,
		answer: 4,
	},
	// Test Case 1: Target not found (between elements)
	{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 2,
		answer: -1,
	},
	// Test Case 2: Target found (single element array)
	{
		nums:   []int{5},
		target: 5,
		answer: 0,
	},
	// Test Case 3: Target not found (single element array)
	{
		nums:   []int{5},
		target: -5,
		answer: -1,
	},
	// Test Case 4: Empty array
	{
		nums:   []int{},
		target: 0,
		answer: -1,
	},
	// Test Case 5: Target found at the end
	{
		nums:   []int{2, 5},
		target: 5,
		answer: 1,
	},
	// Test Case 6: Target found at the beginning
	{
		nums:   []int{2, 5},
		target: 2,
		answer: 0,
	},
	// Test Case 7: Target not found (larger than max)
	{
		nums:   []int{2, 5},
		target: 6,
		answer: -1,
	},
	// Test Case 8: Target not found (smaller than min)
	{
		nums:   []int{2, 5},
		target: 1,
		answer: -1,
	},
	// Test Case 9: Larger array, target present
	{
		nums:   []int{-10, -5, 0, 3, 7, 11, 15, 20},
		target: 15,
		answer: 6,
	},
	// Test Case 10: Larger array, target absent
	{
		nums:   []int{-10, -5, 0, 3, 7, 11, 15, 20},
		target: 1,
		answer: -1,
	},
	// Test Case 11: Array with duplicates (binary search finds one instance)
	{
		nums:   []int{1, 2, 2, 2, 3, 4},
		target: 2,
		answer: 2, // Standard binary search often finds the middle '2'
	},
}
