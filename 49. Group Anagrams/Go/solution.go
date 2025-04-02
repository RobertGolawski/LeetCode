package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	type KeyType [26]int
	res := make(map[KeyType][]string)
	for _, str := range strs {
		strKey := KeyType{}
		for _, c := range str {
			strKey[rune(c)-rune('a')] += 1
		}
		if _, ok := res[strKey]; !ok {
			res[strKey] = []string{}
		}
		res[strKey] = append(res[strKey], str)
	}
	ans := [][]string{}
	for _, value := range res {
		ans = append(ans, value)
	}
	return ans
}

func compare(expected map[string]bool, actual [][]string) bool {
	if len(expected) != len(actual) {
		return false
	}

	actualMap := make(map[string]bool)
	for _, group := range actual {
		sort.Strings(group)
		key := strings.Join(group, ",")
		actualMap[key] = true
	}

	for key := range expected {
		if _, ok := actualMap[key]; !ok {
			return false
		}
	}
	return true
}

func main() {
	for i := range T.cases {
		res := groupAnagrams(T.cases[i])
		ans := T.answers[i]
		passed := compare(ans, res)
		if !passed {
			fmt.Printf("Test Case %d Failed:\n", i)
			fmt.Printf("  Input: %v\n", T.cases[i])
			fmt.Printf("  Expected: %v\n", T.answers[i])
			fmt.Printf("  Got: %v\n", ans)
			return
		}
	}
	fmt.Print("Congratulations")
}
