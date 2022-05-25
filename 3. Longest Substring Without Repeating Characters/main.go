package main

import "fmt"

// Given a string s,
// find the length of the longest substring without repeating characters.
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	charMap := map[int32]int{}

	var max int
	var startID int

	var updateMax = func() {
		if len(charMap) > max {
			max = len(charMap)
		}
	}

	for i, char := range s {
		duplicatedID, ok := charMap[char]

		if ok {
			updateMax()

			for j := startID; j <= duplicatedID; j++ {
				delete(charMap, int32(s[j]))
			}

			startID = duplicatedID + 1
		}

		charMap[char] = i
	}

	updateMax()

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("umvejcuuk"))
	fmt.Println(lengthOfLongestSubstring("ggububgvfk"))
}
