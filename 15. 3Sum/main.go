package main

import (
	"fmt"
	"sort"
)

// Given an array nums of n integers,
// are there elements a, b, c in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
//
// Note:
//
// The solution set must not contain duplicate triplets.

// Given array nums = [-1, 0, 1, 2, -1, -4],
//
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// Memory Limit Exceeded
func threeSum(nums []int) [][]int {
	var result [][]int

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	sumMap := map[int]map[int]map[int]int{}
	duplicateSet := map[string]struct{}{}

	for i, iValue := range nums {
		firstMap, ok := sumMap[iValue]
		if !ok {
			firstMap = map[int]map[int]int{}
			sumMap[iValue] = firstMap
		} else {
			continue
		}

		for j, jValue := range nums {
			if i == j {
				continue
			}

			secondMap, ok := firstMap[jValue]
			if !ok {
				secondMap = map[int]int{}
				firstMap[jValue] = secondMap
			} else {
				continue
			}

			for k := len(nums) - 1; k > j && i != k; k-- {
				kValue := nums[k]

				if _, ok := secondMap[kValue]; !ok {
					sum := iValue + jValue + kValue
					secondMap[kValue] = sum

					if sum != 0 {
						continue
					}

					values := []int{iValue, jValue, kValue}
					sort.Ints(values)
					key := fmt.Sprintf("%d-%d-%d", values[0], values[1], values[2])

					if _, ok := duplicateSet[key]; !ok {
						result = append(result, []int{iValue, jValue, kValue})
						duplicateSet[key] = struct{}{}
					}
				}
			}

		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{1, 2, -2, -1}))
}
