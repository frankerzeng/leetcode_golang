/**
53. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2}))
}

// 两块夹板分隔元素
func maxSubArray(nums []int) int {
	rst := nums[0]
	var tmpSum int
	for i := 0; i < len(nums); i++ {
		tmpSum = nums[i]
		if tmpSum > rst {
			rst = tmpSum
		}
		for j := i + 1; j < len(nums); j++ {
			tmpSum += nums[j]
			if tmpSum > rst {
				rst = tmpSum
			}
		}
	}

	return rst
}

// best solution
// func maxSubArray(nums []int) int {
// 	pre, ret := 0, math.MinInt32
// 	for _, v := range nums {
// 		if pre > 0 {
// 			pre += v
// 		} else {
// 			pre = v
// 		}
// 		if pre > ret { ret = pre }
// 	}
// 	return ret
// }
