package main

import "fmt"

// You are given an integer array nums.
// You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

func canJump(nums []int) bool {
	count := len(nums)

	if count == 0 {
		return false
	}

	max := nums[0]

	for i := 0; i <= max; i++ {
		value := i + nums[i]

		if value >= count-1 {
			return true
		}

		if value > max {
			max = value
		}
	}

	return false
}

func canJumpExample(nums []int) bool {
	lastPos := len(nums) - 1 // last position can reach the end index
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{1}))
}
