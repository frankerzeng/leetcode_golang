/*
442. Find All Duplicates in an Array
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

func findDuplicates(nums []int) []int {
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
