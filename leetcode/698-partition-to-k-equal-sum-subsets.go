/*
698. Partition to K Equal Sum Subsets

Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty
subsets whose sums are all equal.

Example 1:
	Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
	Output: True
	Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
3 2 2 1 k=2
Note:
	1 <= k <= len(nums) <= 16.
	0 < nums[i] < 10000.

*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// nums := []int{5524, 4485, 4391, 3908, 2299, 1601, 1175, 1100, 1024, 730, 659, 580, 405, 401, 383, 3}
	// k := 4

	nums := []int{15, 3557, 42, 3496, 5, 81, 34, 95, 9, 81, 42, 106, 71}
	k := 11
	start := time.Now()
	fmt.Println(canPartitionKSubsets(nums, k))
	fmt.Println(time.Since(start))
}
func canPartitionKSubsets(nums []int, k int) bool {
	start := time.Now()
	fmt.Println(time.Since(start))

	sum := 0
	for _, v := range nums {
		sum += v
	}
	partSumInt := sum / k

	// can't divided by k
	if sum%k != 0 {
		return false
	}

	return tFun(nums, partSumInt, make([]int, k), len(nums)-1)
}

// target sum
// left is the numeric who the sum equal target
func tFun(nums []int, sum int, sumMap []int, index int) bool {
	fmt.Println("++++++++++++++++++", index, sumMap, " |index=", index, " |target=", sum)

	if index == -1 {
		for _, v := range sumMap {
			fmt.Println(v)
			if v != sum {
				return false
			}
		}
		return true
	}
	num := nums[index]
	for i := 0; i < len(sumMap); i++ {
		// fmt.Println(sumMap, i, num, "indexI=", index)
		if sumMap[i]+num <= sum {
			// fmt.Println("---------")
			sumMap[i] = sumMap[i] + num
			if index == 0 {
				// fmt.Println("---------")
				// fmt.Println(nums,sum,sumMap,index-1)
			}
			if tFun(nums, sum, sumMap, index-1) {
				return true
			}
			sumMap[i] = sumMap[i] - num
		}
	}
	return false
}
