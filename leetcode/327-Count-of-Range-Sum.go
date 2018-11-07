/*
327. Count of Range Sum
Given an integer array nums, return the number of range sums that lie in [lower, upper] inclusive.
Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j (i ≤ j), inclusive.

Note:
A naive algorithm of O(n2) is trivial. You MUST do better than that.

Example:

Input: nums = [-2,5,-1], lower = -2, upper = 2,
Output: 3
Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective sums are: -2, -1, 2.
*/
package main

import "fmt"

func main() {
	nums := []int{-2, 5, -1}
	fmt.Println(countRangeSum(nums, -2, 2))
}

// method 1 ,Time Limit Exceeded
func countRangeSum1(nums []int, lower int, upper int) int {
	rst := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			tmp := 0
			for k := i; k <= j; k++ { // s(i,j)
				tmp += nums[k]
			}
			if tmp <= upper && tmp >= lower {
				rst++
			}
		}
	}
	return rst
}
func countRangeSum(nums []int, lower int, upper int) int {
	rst := 0
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		tmp := 0
		for j := i; j < lenNums; j++ {
			if i == j {
				tmp = nums[i]
			} else {
				tmp = tmp + nums[j]
			}
			if tmp <= upper && tmp >= lower {
				rst++
			}

		}
	}

	return rst
}
