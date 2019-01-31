/*
561. Array Partition I
Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ...,
(an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
Note:
	1.n is a positive integer, which is in the range of [1, 10000].
	2.All the integers in the array will be in the range of [-10000, 10000].
*/
package main

import "fmt"

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

func arrayPairSum(nums []int) int {
	var rst int
	// 快排
	quickSortFunc(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i = i + 2 {
		rst += nums[i]
	}
	return rst
}

func quickSortFunc(values []int, left int, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSortFunc(values, left, p-1)
	}
	if right-p > 1 {
		quickSortFunc(values, p+1, right)
	}
}
