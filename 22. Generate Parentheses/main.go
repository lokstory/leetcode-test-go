package main

import (
	"fmt"
)

// Given n pairs of parentheses,
// write a function to generate all combinations of well-formed parentheses.

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

func generateParenthesis(n int) []string {
	var ret []string
	backtrack(n, n, &ret, "")
	return ret
}

func backtrack(left, right int, ret *[]string, cur string) {
	if left == 0 && right == 0 {
		*ret = append(*ret, cur)
		return
	}

	if right < left {
		return
	}

	if left > 0 {
		backtrack(left-1, right, ret, cur+"(")
	}

	if right > 0 {
		backtrack(left, right-1, ret, cur+")")
	}
}

func generateParenthesis2(n int) []string {
	var ret []string

	if n == 0 {
		ret = append(ret, "")
	} else {
		for i := 0; i < n; i++ {
			for _, left := range generateParenthesis2(i) {
				for _, right := range generateParenthesis2(n-1-i) {
					ret = append(ret, fmt.Sprintf("(%s)%s", left, right))
				}
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(generateParenthesis(3))
	//fmt.Println(generateParenthesis2(3))
}
