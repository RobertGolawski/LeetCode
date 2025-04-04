package main

type tests struct {
	cases   []string
	answers []bool
}

var T = tests{
	cases: []string{
		// --- Basic Valid Cases ---
		"()",     // Simple pair
		"()[]{}", // Multiple pairs
		"{[]}",   // Nested pairs

		// --- Basic Invalid Cases ---
		"(]",   // Mismatched pair
		"([)]", // Incorrect nesting order
		"}",    // Single closing bracket
		"{",    // Single opening bracket

		// --- Edge Cases ---
		"",    // Empty string (considered valid)
		"((",  // Only opening brackets
		"))",  // Only closing brackets
		"({[", // Unclosed opening brackets
		"]})", // Unopened closing brackets

		// --- More Complex Cases ---
		"({[]})",            // Valid complex nesting
		"((()))",            // Valid deep nesting
		"([]{})",            // Valid mixed nesting
		"[(])",              // Invalid nesting
		"{[}]",              // Invalid nesting
		"())",               // Invalid - premature closing
		"(()",               // Invalid - unclosed opening
		"((()))[]{}{[()]}",  // Long valid string
		"((()))[]{}{[()]}}", // Long invalid string (extra closer)
		"](){}",             // Invalid - starts with closer
		"(){}[",             // Invalid - ends with opener
	},
	answers: []bool{
		// --- Basic Valid Cases ---
		true, // "()"
		true, // "()[]{}"
		true, // "{[]}"

		// --- Basic Invalid Cases ---
		false, // "(]"
		false, // "([)]"
		false, // "}"
		false, // "{"

		// --- Edge Cases ---
		true,  // ""
		false, // "(("
		false, // "))"
		false, // "({["
		false, // "]})"

		// --- More Complex Cases ---
		true,  // "({[]})"
		true,  // "((()))"
		true,  // "([]{})"
		false, // "[(])"
		false, // "{[}]"
		false, // "())"
		false, // "(()"
		true,  // "((()))[]{}{[()]}",
		false, // "((()))[]{}{[()]}}"
		false, // "](){}"
		false, // "(){}[",
	},
}
