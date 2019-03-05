/*
321. Create Maximum Number
Given two arrays of length m and n with digits 0-9 representing two numbers. Create the maximum number of length k <= m + n from digits of the two. The relative order of the digits from the same array must be preserved. Return an array of the k digits.

Note: You should try to optimize your time and space complexity.

Example 1:

Input:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
Output:
[9, 8, 6, 5, 3]
Example 2:

Input:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
Output:
[6, 7, 6, 0, 4]
Example 3:

Input:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
Output:
[9, 8, 9]
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxNumber([]int{8, 5, 9, 5, 1, 6, 9}, []int{2, 6, 4, 3, 8, 4, 1, 0, 7, 2, 9, 2, 8}, 20))
	// fmt.Println(maxNumber([]int{3, 6, 5}, []int{3}, 3))
	// [9, 8, 6, 5, 3]
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var rst []int
	maxNum := maxNumberFunc(nums1, nums2, k)
	for k > 0 {
		rst = append(rst, maxNum/int(math.Pow(10, float64(k-1))))
		maxNum = maxNum % int(math.Pow(10, float64(k-1)))
		k--
	}
	return rst
}

// 以3开头的
func maxNumberFunc(nums1 []int, nums2 []int, k int) int {
	if len(nums1)+len(nums2) < k {
		panic("sdfs")
	}
	max := 0
	tmp := 0
	for index, val := range nums1 {
		if k == 1 {
			if max < val {
				max = val * int(math.Pow(10, float64(k-1)))
			}
		} else {
			if (len(nums1)+len(nums2)-index) >= k && (len(nums1)-index) > 0 {
				rst := maxNumberFunc(nums1[index+1:], nums2, k-1)
				if tmp < rst+val*int(math.Pow(10, float64(k-1))) {
					tmp = rst + val*int(math.Pow(10, float64(k-1)))
				}
			}
		}
	}
	if max < tmp {
		max = tmp
	}

	max1 := 0
	tmp = 0
	for index, val := range nums2 {
		if k == 1 {
			if max1 < val {
				max1 = val * int(math.Pow(10, float64(k-1)))
			}
		} else {
			if (len(nums1)+len(nums2)-index) >= k && (len(nums2)-index) > 0 {
				rst := maxNumberFunc(nums1, nums2[index+1:], k-1)
				if tmp < rst+val*int(math.Pow(10, float64(k-1))) {
					tmp = rst + val*int(math.Pow(10, float64(k-1)))
				}
			}
		}
	}
	if max1 < tmp {
		max1 = tmp
	}

	if max1 > max {
		max = max1
	}
	return max
}
