/*
347. Top K Frequent Elements
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

*/
package main

import "fmt"

func main() {
	nums := []int{1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	var rst []int

	numsMap := map[int]int{}
	for _, v := range nums {
		numsMap[v]++
	}

	numsMapSort := map[int][]int{}     // map [times=>[num1,num2]]
	maxTmp := make([]int, len(nums)+1) // list of frequent times
	for k, v := range numsMap {
		numsMapSort[v] = append(numsMapSort[v], k)
		maxTmp[v] = 1
	}

	for i := len(maxTmp) - 1; i >= 0; i-- {
		if maxTmp[i] != 0 {
			for _, vv := range numsMapSort[i] {
				if k == 0 {
					return rst
				}
				rst = append(rst, vv)
				k--
			}
		}
	}

	return rst
}
