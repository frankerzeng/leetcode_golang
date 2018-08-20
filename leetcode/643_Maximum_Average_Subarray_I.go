/**
643. Maximum Average Subarray I
Given an array consisting of n integers, find the contiguous subarray of given length k that has the maximum average value.
And you need to output the maximum average value.

Example 1:
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
Note:
1 <= k <= n <= 30,000.
Elements of the given array will be in the range [-10,000, 10,000].
*/
package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
func findMaxAverage(nums []int, k int) float64 {
	rst := 0
	for i := 0; i < k; i++ {
		rst += nums[i]
	}

	for i := k; i < len(nums); i++ {
		rstTmp := 0
		// sum of k cell
		for j := i; j > i-k; j-- {
			rstTmp += nums[j]
		}
		if rstTmp > rst {
			rst = rstTmp
		}
	}
	return float64(rst) / float64(k)
}
