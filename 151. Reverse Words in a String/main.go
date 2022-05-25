package main

import (
	"fmt"
	"strings"
)

// Given an input string, reverse the string word by word.
// Example 1:
// Input: "the sky is blue"
// Output: "blue is sky the"
// Example 2:
// Input: "  hello world!  "
// Output: "world! hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:
// Input: "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
// Note:
// A word is defined as a sequence of non-space characters.
// Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
// You need to reduce multiple spaces between two words to a single space in the reversed string.
// Follow up:
// For C programmers, try to solve it in-place in O(1) extra space.

func reverseWords(s string) string {
	fields := strings.Fields(s)

	var reverses []string
	for i := len(fields) - 1; i >= 0; i-- {
		reverses = append(reverses, fields[i])
	}

	return strings.Join(reverses, " ")
}

func main() {
	fmt.Println(reverseWords("lak.kb c!gfbb' cgyxxrph!ai paim,izbj.tnkugjx.f!uhs!xgv vsx.ncydmsgeaenstgthzd'fv qssjheigf!xca!d ,tsvj!yni'csdnphtt cej.ngxy egnh oaxzxugnehorkqkt,"))
}
