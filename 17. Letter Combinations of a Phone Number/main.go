package main

import (
	"fmt"
)

// Given a string containing digits from 2-9 inclusive,
// return all possible letter combinations that the number could represent.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.

// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

var charMap = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var result []string

	if len(digits) == 0 {
		return result
	}

	for _, key := range digits {
		list, ok := charMap[key]
		if !ok {
			continue
		}

		if len(result) == 0 {
			for _, c := range list {
				result = append(result, string(c))
			}
			continue
		}

		var newList []string
		for _, text := range result {
			for _, c := range list {
				newList = append(newList, text + string(c))
			}
		}
		result = newList
	}

	return result
}

func recursiveSolution(digits string) []string {
	var result []string

	if len(digits) == 0 {
		return result
	}

	var backtracking func(current string, nextDigits string)

	backtracking = func(current string, nextDigits string) {
		if len(nextDigits) == 0 {
			result = append(result, current)
			return
		}

		chars, ok := charMap[rune(nextDigits[0])]
		if !ok {
			return
		}

		for _, letter := range chars {
			backtracking(current + string(letter), nextDigits[1:])
		}
	}


	backtracking("", digits)
	return result
}


func main() {
	//fmt.Println(letterCombinations("23"))
	fmt.Println(recursiveSolution("23"))
}
