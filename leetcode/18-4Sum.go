/**
18. 4Sum
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that
a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
-2 -1 0 0 1 2
A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	existMap = make(map[string]int)
	return xSum(nums, target, 4)
}

var existMap map[string]int

func xSum(nums []int, target int, step int) [][]int {
	var rst [][]int
	if step == 1 {
		for _, v := range nums {
			if v == target {
				rst = append(rst, []int{v})
			}
		}
	} else {
		for i := 0; i <= len(nums)-step; i++ {
			for _, v := range xSum(nums[i+1:], target-nums[i], step-1) {
				tmpSlice := append([]int{nums[i]}, v...)
				if step == 4 {
					key := strconv.Itoa(tmpSlice[0]) + "+" + strconv.Itoa(tmpSlice[1]) + "+" + strconv.Itoa(tmpSlice[2]) + "+" + strconv.Itoa(tmpSlice[3])
					if _, ok := existMap[key]; !ok {
						existMap[key] = 1
					} else {
						continue
					}
				}
				rst = append(rst, tmpSlice)
			}
		}
	}

	return rst
}
