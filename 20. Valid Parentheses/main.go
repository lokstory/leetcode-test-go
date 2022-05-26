package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
//
// Input: s = "()"
// Output: true
//
// Input: s = "()[]{}"
// Output: true
//
// Input: s = "(]"
// Output: false
//
// Constraints:
// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.

var bracketMap = map[int32]int32{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	count := len(s)

	if count == 0 || count%2 == 1 {
		return false
	}

	var leftBrackets []int32
	var leftCount int

	for _, c := range s {
		leftBracket, ok := bracketMap[c]

		if !ok {
			leftBrackets = append(leftBrackets, c)
			continue
		}

		leftCount = len(leftBrackets)

		if leftCount == 0 || leftBrackets[leftCount-1] != leftBracket {
			return false
		}

		leftBrackets = leftBrackets[:leftCount-1]
	}

	return len(leftBrackets) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
}
