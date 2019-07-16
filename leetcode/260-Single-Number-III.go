/**
260. Single Number III
Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly
twice. Find the two elements that appear only once.

Example:
Input:  [1,2,1,3,2,5]
Output: [3,5]
Note:
	1.The order of the result is not important. So in the above example, [5, 3] is also correct.
	2.Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?
*/
package main

import "fmt"

func main() {
	// singleNumber([]int{1, 2, 2, 3})
	// singleNumber([]int{1,2,1,3,2,5})
	fmt.Println(singleNumber([]int{0, 1, 2, 2}))
}
func singleNumber(nums []int) []int {
	index := 0
I:
	for true {
		if index >= len(nums)-1 {
			break
		}
		for i := index + 1; i < len(nums); i++ {
			if nums[i] == nums[index] {
				nums = append(append(nums[0:index], nums[index+1:i]...), nums[i+1:]...)
				continue I
			}
		}
		index++
	}
	return nums
}
