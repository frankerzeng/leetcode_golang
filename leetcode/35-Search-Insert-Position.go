/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it
would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/
package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1}, 2))
}
func searchInsert(nums []int, target int) int {
	var rst int

	left := 0
	right := len(nums) - 1
	middle := (left + right) / 2

	if right == -1 {
		return 0
	}

	// binary search
	for {
		if nums[right] < target {
			return right + 1
		}
		if nums[left] > target {
			return 0
		}
		if nums[right] == target {
			return right
		}
		if nums[left] == target {
			return left
		}

		if (left + 1) == right {
			return right
		}

		// 中间位置
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle
			middle = (left + right) / 2
		} else {
			left = middle
			middle = (left + right) / 2
		}
	}

	return rst
}
