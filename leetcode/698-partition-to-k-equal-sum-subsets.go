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
)

func main() {
	nums := []int{15, 11, 10, 4, 3, 1, 1, 1, 1, 1}
	k := 5
	fmt.Println(canPartitionKSubsets(nums, k))
}
func canPartitionKSubsets(nums []int, k int) bool {
	lenNumb := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	partSum := float32(sum) / float32(k)
	partSumInt := sum / k

	// can't divided by k
	if partSum != float32(partSumInt) {
		return false
	}

	/*max item numbers of one part
	example: [4, 3, 2, 3, 5, 2, 1], k = 4 possibly subset (4)(3)(2)(3,5,2,1)
	then max part is (3,5,2,1)
		[5,4, 3, 3, 3, 2, 2, 1], k = 4 possibly subset (4)(3)(2)(3,5,2,1)
	*/
	quickSortDesc(nums, 0, lenNumb-1)
	if tFun(nums, partSumInt, partSumInt, k) {
		return true
	}

	return false
}

// target sum
// left is the numeric who the sum equal target
func tFun(nums []int, left int, target int, k int) bool {
	if len(nums) == 1 {
		if nums[0] == left && k == 1 {
			return true
		} else {
			return false
		}
	} else {
		if nums[0] == left && tFun(nums[1:], left-nums[0], target, k) {
			return true
		} else if nums[0] > left {
			return tFun(nums[0])
		}
		{
			return
		}
	}
	return false
}
func quickSortDesc(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] <= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] >= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSortDesc(values, left, p-1)
	}
	if right-p > 1 {
		quickSortDesc(values, p+1, right)
	}
}
