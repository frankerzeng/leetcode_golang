/*
442. Find All Duplicates in an Array
找出数组内重复的元素

Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDuplicates([]int{3, 3}))
}

// 方法一 faster than 10.49%
func findDuplicates1(nums []int) []int {
	sort.Ints(nums)

	var rst []int
	for i := 0; i < len(nums); i++ {
		if i != len(nums)-1 {
			if nums[i] == nums[i+1] {
				rst = append(rst, nums[i])
				for i != len(nums)-1 && nums[i] == nums[i+1] {
					i++
				}
			}
		}
	}

	return rst
}
func findDuplicates(nums []int) []int {
	var rstMap map[int]int
	rstMap = make(map[int]int)
	var rst []int
	for _, v := range nums {
		if rstMap[v] == 0 {
			rstMap[v] = 1
		} else if rstMap[v] == 1 {
			rst = append(rst, v)
			rstMap[v] = 2
		}
	}

	return rst
}
