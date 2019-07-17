/**
239. Sliding Window Maximum
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very
right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return
the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

Follow up:
Could you solve it in linear time?
*/

package main

func main() {
	maxSlidingWindow([]int{}, 0)
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}
func maxSlidingWindow(nums []int, k int) []int {
	var rstSlice []int
	if len(nums) == 0 || k == 0 {
		return rstSlice
	}

	currentSlice := nums[0:k]              // first slide
	maxIndex := maxIndexFunc(currentSlice) // max value's index
	maxValue := currentSlice[maxIndex]     // max value

	rstSlice = append(rstSlice, maxValue)

	for i := k; i < len(nums); i++ {
		currentSlice = nums[i+1-k : i+1]
		if maxIndex == 0 {
			maxIndex = maxIndexFunc(currentSlice)
			maxValue = currentSlice[maxIndex]
		} else { // pre max value also in next sliding window
			if maxValue > nums[i] {
				maxIndex--
				maxValue = currentSlice[maxIndex]
			} else {
				maxIndex = k - 1
				maxValue = currentSlice[maxIndex]
			}
		}
		rstSlice = append(rstSlice, maxValue)
	}

	return rstSlice
}

// find the max value's index of slice
func maxIndexFunc(nums []int) int {
	maxVal, maxIndex := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex = i
		}
	}
	return maxIndex
}
