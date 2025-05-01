package main

// testCase defines a sequence of operations and expected results for MinStack
type testCase struct {
	commands  []string      // Sequence of method names ("Constructor", "Push", "Pop", "Top", "GetMin")
	arguments [][]int       // Arguments for each command (e.g., {val} for Push, {} for others)
	expected  []interface{} // Expected return values (int for Top/GetMin, nil for others)
}

// T holds all the test cases
var T = []testCase{
	// Test Case 0: Basic operations
	{
		commands:  []string{"Constructor", "Push", "Push", "Push", "GetMin", "Pop", "Top", "GetMin"},
		arguments: [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
		expected:  []interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
	},
	// Test Case 1: Popping the minimum
	{
		commands:  []string{"Constructor", "Push", "Push", "GetMin", "Pop", "GetMin"},
		arguments: [][]int{{}, {0}, {-1}, {}, {}, {}},
		expected:  []interface{}{nil, nil, nil, -1, nil, 0},
	},
	// Test Case 2: Empty stack operations
	{
		commands:  []string{"Constructor", "Pop", "Top", "GetMin"},
		arguments: [][]int{{}, {}, {}, {}},
		expected:  []interface{}{nil, nil, 0, 0}, // Assuming 0 for Top/GetMin on empty stack as per your code
	},
	// Test Case 3: More complex sequence
	{
		commands:  []string{"Constructor", "Push", "Push", "Push", "Top", "Pop", "GetMin", "Pop", "GetMin", "Pop", "GetMin"},
		arguments: [][]int{{}, {2}, {0}, {3}, {}, {}, {}, {}, {}, {}, {}},
		expected:  []interface{}{nil, nil, nil, nil, 3, nil, 0, nil, 2, nil, 0},
	},
	// Test Case 4: Handling duplicates of the minimum
	{
		commands:  []string{"Constructor", "Push", "Push", "Push", "GetMin", "Pop", "GetMin", "Pop", "GetMin"},
		arguments: [][]int{{}, {-1}, {-1}, {0}, {}, {}, {}, {}, {}},
		expected:  []interface{}{nil, nil, nil, nil, -1, nil, -1, nil, -1},
	},
}

// Note: 'nil' in expected means we don't check a return value for that specific step (like Push, Pop, Constructor).
// We check the state implicitly via subsequent Top/GetMin calls.
