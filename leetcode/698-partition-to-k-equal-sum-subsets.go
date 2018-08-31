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
	// nums := []int{5524, 4485, 4391, 3908, 2299, 1601, 1175, 1100, 1024, 730, 659, 580, 405, 401, 383, 3}
	// k := 4

	nums := []int{4, 5, 3, 2, 5, 5, 5, 1, 5, 5, 5, 5, 5, 5, 5, 5}
	k := 14
	fmt.Println(canPartitionKSubsets(nums, k))
}
func canPartitionKSubsets(nums []int, k int) bool {
	lenNumb := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	partSumInt := sum / k

	// can't divided by k
	if sum%k != 0 {
		return false
	}

	/*max item numbers of one part
	example: [4, 3, 2, 3, 5, 2, 1], k = 4 possibly subset (4)(3)(2)(3,5,2,1)
	then max part is (3,5,2,1)
		[5,4, 3, 3, 3, 2, 2, 1], k = 4 possibly subset (4)(3)(2)(3,5,2,1)
	*/
	quickSortDesc(nums, 0, lenNumb-1)
	return tFun(nums, partSumInt, make([]int, k), len(nums)-1)
}

var indexI int

// target sum
// left is the numeric who the sum equal target
func tFun(nums []int, sum int, sums []int, index int) bool {
	indexI++
	if indexI > 20 {
		return false
	}

	fmt.Println(sums)

	if index == -1 {
		for _, v := range sums {
			if v != sum {
				return false
			}
		}
		return true
	}
	num := nums[index]
	for i := 0; i < len(sums); i++ {
		if sums[i]+num <= sum {
			sums[i] += num
			if tFun(nums, sum, sums, index-1) {
				return true
			}
			sums[i] -= num
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
