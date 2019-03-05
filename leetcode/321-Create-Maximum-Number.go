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

	// fmt.Println(maxNumber2(
	// 	[]int{8, 9, 7, 3, 5, 9, 1, 0, 8, 5, 3, 0, 9, 2, 7, 4, 8, 9, 8, 1, 0, 2, 0, 2, 7, 2, 3, 5, 4, 7, 4, 1, 4, 0, 1, 4, 2, 1, 3, 1, 5, 3, 9, 3, 9, 0, 1, 7, 0, 6, 1, 8, 5, 6, 6, 5, 0, 4, 7, 2, 9, 2, 2, 7, 6, 2, 9, 2, 3, 5, 7, 4, 7, 0, 1, 8, 3, 6, 6, 3, 0, 8, 5, 3, 0, 3, 7, 3, 0, 9, 8, 5, 1, 9, 5, 0, 7, 9, 6, 8, 5, 1, 9, 6, 5, 8, 2, 3, 7, 1, 0, 1, 4, 3, 4, 4, 2, 4, 0, 8, 4, 6, 5, 5, 7, 6, 9, 0, 8, 4, 6, 1, 6, 7, 2, 0, 1, 1, 8, 2, 6, 4, 0, 5, 5, 2, 6, 1, 6, 4, 7, 1, 7, 2, 2, 9, 8, 9, 1, 0, 5, 5, 9, 7, 7, 8, 8, 3, 3, 8, 9, 3, 7, 5, 3, 6, 1, 0, 1, 0, 9, 3, 7, 8, 4, 0, 3, 5, 8, 1, 0, 5, 7, 2, 8, 4, 9, 5, 6, 8, 1, 1, 8, 7, 3, 2, 3, 4, 8, 7, 9, 9, 7, 8, 5, 2, 2, 7, 1, 9, 1, 5, 5, 1, 3, 5, 9, 0, 5, 2, 9, 4, 2, 8, 7, 3, 9, 4, 7, 4, 8, 7, 5, 0, 9, 9, 7, 9, 3, 8, 0, 9, 5, 3, 0, 0, 3, 0, 4, 9, 0, 9, 1, 6, 0, 2, 0, 5, 2, 2, 6, 0, 0, 9, 6, 3, 4, 1, 2, 0, 8, 3, 6, 6, 9, 0, 2, 1, 6, 9, 2, 4, 9, 0, 8, 3, 9, 0, 5, 4, 5, 4, 6, 1, 2, 5, 2, 2, 1, 7, 3, 8, 1, 1, 6, 8, 8, 1, 8, 5, 6, 1, 3, 0, 1, 3, 5, 6, 5, 0, 6, 4, 2, 8, 6, 0, 3, 7, 9, 5, 5, 9, 8, 0, 4, 8, 6, 0, 8, 6, 6, 1, 6, 2, 7, 1, 0, 2, 2, 4, 0, 0, 0, 4, 6, 5, 5, 4, 0, 1, 5, 8, 3, 2, 0, 9, 7, 6, 2, 6, 9, 9, 9, 7, 1, 4, 6, 2, 8, 2, 5, 3, 4, 5, 2, 4, 4, 4, 7, 2, 2, 5, 3, 2, 8, 2, 2, 4, 9, 8, 0, 9, 8, 7, 6, 2, 6, 7, 5, 4, 7, 5, 1, 0, 5, 7, 8, 7, 7, 8, 9, 7, 0, 3, 7, 7, 4, 7, 2, 0, 4, 1, 1, 9, 1, 7, 5, 0, 5, 6, 6, 1, 0, 6, 9, 4, 2, 8, 0, 5, 1, 9, 8, 4, 0, 3, 1, 2, 4, 2, 1, 8, 9, 5, 9, 6, 5, 3, 1, 8, 9, 0, 9, 8, 3, 0, 9, 4, 1, 1, 6, 0, 5, 9, 0, 8, 3, 7, 8, 5},
	// 	[]int{7, 8, 4, 1, 9, 4, 2, 6, 5, 2, 1, 2, 8, 9, 3, 9, 9, 5, 4, 4, 2, 9, 2, 0, 5, 9, 4, 2, 1, 7, 2, 5, 1, 2, 0, 0, 5, 3, 1, 1, 7, 2, 3, 3, 2, 8, 2, 0, 1, 4, 5, 1, 0, 0, 7, 7, 9, 6, 3, 8, 0, 1, 5, 8, 3, 2, 3, 6, 4, 2, 6, 3, 6, 7, 6, 6, 9, 5, 4, 3, 2, 7, 6, 3, 1, 8, 7, 5, 7, 8, 1, 6, 0, 7, 3, 0, 4, 4, 4, 9, 6, 3, 1, 0, 3, 7, 3, 6, 1, 0, 0, 2, 5, 7, 2, 9, 6, 6, 2, 6, 8, 1, 9, 7, 8, 8, 9, 5, 1, 1, 4, 2, 0, 1, 3, 6, 7, 8, 7, 0, 5, 6, 0, 1, 7, 9, 6, 4, 8, 6, 7, 0, 2, 3, 2, 7, 6, 0, 5, 0, 9, 0, 3, 3, 8, 5, 0, 9, 3, 8, 0, 1, 3, 1, 8, 1, 8, 1, 1, 7, 5, 7, 4, 1, 0, 0, 0, 8, 9, 5, 7, 8, 9, 2, 8, 3, 0, 3, 4, 9, 8, 1, 7, 2, 3, 8, 3, 5, 3, 1, 4, 7, 7, 5, 4, 9, 2, 6, 2, 6, 4, 0, 0, 2, 8, 3, 3, 0, 9, 1, 6, 8, 3, 1, 7, 0, 7, 1, 5, 8, 3, 2, 5, 1, 1, 0, 3, 1, 4, 6, 3, 6, 2, 8, 6, 7, 2, 9, 5, 9, 1, 6, 0, 5, 4, 8, 6, 6, 9, 4, 0, 5, 8, 7, 0, 8, 9, 7, 3, 9, 0, 1, 0, 6, 2, 7, 3, 3, 2, 3, 3, 6, 3, 0, 8, 0, 0, 5, 2, 1, 0, 7, 5, 0, 3, 2, 6, 0, 5, 4, 9, 6, 7, 1, 0, 4, 0, 9, 6, 8, 3, 1, 2, 5, 0, 1, 0, 6, 8, 6, 6, 8, 8, 2, 4, 5, 0, 0, 8, 0, 5, 6, 2, 2, 5, 6, 3, 7, 7, 8, 4, 8, 4, 8, 9, 1, 6, 8, 9, 9, 0, 4, 0, 5, 5, 4, 9, 6, 7, 7, 9, 0, 5, 0, 9, 2, 5, 2, 9, 8, 9, 7, 6, 8, 6, 9, 2, 9, 1, 6, 0, 2, 7, 4, 4, 5, 3, 4, 5, 5, 5, 0, 8, 1, 3, 8, 3, 0, 8, 5, 7, 6, 8, 7, 8, 9, 7, 0, 8, 4, 0, 7, 0, 9, 5, 8, 2, 0, 8, 7, 0, 3, 1, 8, 1, 7, 1, 6, 9, 7, 9, 7, 2, 6, 3, 0, 5, 3, 6, 0, 5, 9, 3, 9, 1, 1, 0, 0, 8, 1, 4, 3, 0, 4, 3, 7, 7, 7, 4, 6, 4, 0, 0, 5, 7, 3, 2, 8, 5, 1, 4, 5, 8, 5, 6, 7, 5, 7, 3, 3, 9, 6, 8, 1, 5, 1, 1, 1, 0, 3},
	// 	500))
	// [9, 8, 6, 5, 3]
}

// 	Time Limit Exceeded
// 查 nums1 和 nums2 的前某几位的最大值
func maxNumber2(nums1 []int, nums2 []int, k int) []int {
	var rst []int
	str := maxNumberSubstring(nums1, nums2, k)
	for _, v := range str {
		i, _ := strconv.Atoi(string(v))
		rst = append(rst, i)
	}
	return rst
}
func maxNumberSubstring(nums1 []int, nums2 []int, k int) string {
	rst := ""
	tmp1 := -1
	ii := 0
	for i := 0; i < len(nums1) && i <= len(nums1)+len(nums2)-k; i++ {
		if tmp1 < nums1[i] {
			tmp1 = nums1[i]
			ii = i
		}
	}
	tmp2 := -1
	jj := 0
	for j := 0; j < len(nums2) && j <= len(nums1)+len(nums2)-k; j++ {
		if tmp2 < nums2[j] {
			tmp2 = nums2[j]
			jj = j
		}
	}

	if k == 499 {

		fmt.Println(ii, jj)
		// panic(1)
	}

	if k == 1 {
		if tmp1 > tmp2 {
			rst = strconv.Itoa(tmp1)
		} else {
			rst = strconv.Itoa(tmp2)
		}
	} else {
		if tmp1 > tmp2 {
			rst = strconv.Itoa(tmp1) + maxNumberSubstring(nums1[ii+1:], nums2, k-1)
		} else if tmp1 == tmp2 {
			tmpRst1 := maxNumberSubstring(nums1[ii+1:], nums2, k-1)
			tmpRst2 := maxNumberSubstring(nums1, nums2[jj+1:], k-1)
			if tmpRst1 > tmpRst2 {
				rst = strconv.Itoa(tmp1) + tmpRst1
			} else {
				rst = strconv.Itoa(tmp1) + tmpRst2
			}
		} else {
			rst = strconv.Itoa(tmp2) + maxNumberSubstring(nums1, nums2[jj+1:], k-1)
		}
	}
	return rst
}

// 	Time Limit Exceeded
func maxNumber1(nums1 []int, nums2 []int, k int) []int {
	var rst []int
	maxNum := maxNumberFunc(nums1, nums2, k)
	for _, v := range maxNum {
		i, _ := strconv.Atoi(string(v))
		rst = append(rst, i)
	}
	return rst
}

// 	DFS 查子数组的最大值
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
