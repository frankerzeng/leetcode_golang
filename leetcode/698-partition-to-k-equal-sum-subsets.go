/*
698. Partition to K Equal Sum Subsets

Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty
subsets whose sums are all equal.

Example 1:
	Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
	Output: True
	Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.

Note:
	1 <= k <= len(nums) <= 16.
	0 < nums[i] < 10000.

*/
package main

import "fmt"

func main() {
	nums := []int{1, 2}
	k := 1
	fmt.Println(canPartitionKSubsets(nums, k))
}
func canPartitionKSubsets(nums []int, k int) bool {
	ans := false
	//dev
	return ans
}
