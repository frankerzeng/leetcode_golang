/*
485. Max Consecutive Ones
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
*/
package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	rst := 0

	tmp := 0
	for k := range nums {
		if nums[k] == 1 {
			tmp++
			if tmp > rst {
				rst = tmp
			}
		} else {
			tmp = 0
		}

	}

	return rst
}
