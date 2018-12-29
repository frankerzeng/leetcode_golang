/*
229. Majority Element II
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:

Input: [3,2,3]
Output: [3]
Example 2:

Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
*/
package main

func main() {
	majorityElement([]int{1, 2, 3, 3})
}

func majorityElement(nums []int) []int {
	var rst []int

	var rstMap = make(map[int]int)
	for _, v := range nums {
		rstMap[v]++
	}

	for k, v := range rstMap {
		if float32(v) > float32(len(nums)/3) {
			rst = append(rst, k)
		}
	}
	return rst
}
