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
	"strconv"
	"time"
)

func main() {
	si := time.Now()
	// fmt.Println(maxNumber([]int{5, 0, 2, 1, 0, 1, 0, 3, 9, 1, 2, 8, 0, 9, 8, 1, 4, 7, 3}, []int{7, 6, 7, 1, 0, 1, 0, 5, 6, 0, 5, 0}, 31))
	fmt.Println(time.Since(si))
	fmt.Println(maxNumber([]int{8, 9}, []int{3, 9}, 3))
	// [9, 8, 6, 5, 3]
}

// 	Time Limit Exceeded
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var rst []int
	maxNum := maxNumberFunc(nums1, nums2, k)
	for _, v := range maxNum {
		i, _ := strconv.Atoi(string(v))
		rst = append(rst, i)
	}
	return rst
}

// 	Time Limit Exceeded
func maxNumberFunc(nums1 []int, nums2 []int, k int) string {
	if len(nums1)+len(nums2) < k {
		panic("sdfs")
	}
	max := ""
	tmp := ""
	headMax := -1
	for index, val := range nums1 {
		if headMax <= val && (len(nums1)+len(nums2)-index) >= k {
			headMax = val
			if k == 1 {
				if max < strconv.Itoa(val) {
					max = strconv.Itoa(val)
				}
			} else {
				if (len(nums1)+len(nums2)-index) >= k && (len(nums1)-index) > 0 {
					rst := maxNumberFunc(nums1[index+1:], nums2, k-1)
					if tmp < strconv.Itoa(val)+rst {
						tmp = strconv.Itoa(val) + rst
					}
				} else {
					break
				}
			}
		}
	}
	if max < tmp {
		max = tmp
	}

	max1 := ""
	tmp = ""
	for index, val := range nums2 {
		if headMax <= val && (len(nums1)+len(nums2)-index) >= k {
			headMax = val
			if k == 1 {
				if max1 < strconv.Itoa(val) {
					max1 = strconv.Itoa(val)
				}
			} else {
				if (len(nums1)+len(nums2)-index) >= k && (len(nums2)-index) > 0 {
					rst := maxNumberFunc(nums1, nums2[index+1:], k-1)
					if tmp < strconv.Itoa(val)+rst {
						tmp = strconv.Itoa(val) + rst
					}
				} else {
					break
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
