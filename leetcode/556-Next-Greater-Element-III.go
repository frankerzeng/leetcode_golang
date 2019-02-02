/*
556. Next Greater Element III
Given a positive 32-bit integer n, you need to find the smallest 32-bit integer which has exactly the same digits
existing in the integer n and is greater in value than n. If no such positive 32-bit integer exists, you need to return -1.

Example 1:

Input: 12
Output: 21


Example 2:

Input: 21
Output: -1
*/
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(nextGreaterElement(1999999999))
}

// inspired by https://leetcode.com/problems/next-greater-element-iii/discuss/101824/Simple-Java-solution-(4ms)-with-explanation.
func nextGreaterElement(n int) int {
	bit32Int := int(math.Pow(2, 31)) // 不能超过32-bit 的最大数
	if n >= bit32Int {
		return -1
	}
	var rst int
	var nSlice []string
	nStr := strconv.Itoa(n)
	for i := 0; i < len(nStr); i++ {
		nSlice = append(nSlice, nStr[i:i+1])
	}

	var i, j int

	for i = len(nSlice) - 1; i > 0; i-- {
		if nSlice[i-1] < nSlice[i] {
			break
		}
	}

	if i == 0 { // n like 4321 dont has largest num, return -1
		return -1
	}

	x := nSlice[i-1]
	smallIndex := i // 找出非递减的最后一个数
	for j = i + 1; j < len(nSlice); j++ {
		if nSlice[j] > x && nSlice[j] <= nSlice[smallIndex] {
			smallIndex = j
		}
	}

	nSlice[i-1], nSlice[smallIndex] = nSlice[smallIndex], nSlice[i-1] // 非递减的最后的数与最小的数交换位置

	// 非递减的最后的数字从小到大排序
	var nSliceSub []int
	for _, v := range nSlice[i:] {
		tmp, _ := strconv.Atoi(v)
		nSliceSub = append(nSliceSub, tmp)
	}
	sort.Ints(nSliceSub)

	rstSlice := nSlice[0:i]
	for _, v := range nSliceSub {
		rstSlice = append(rstSlice, strconv.Itoa(v))
	}
	for k, v := range rstSlice {
		tmp, _ := strconv.Atoi(v)
		tmpInt := len(rstSlice) - k - 1
		rst += tmp * int(math.Pow(10, float64(tmpInt)))
	}

	if rst >= bit32Int {
		return -1
	}
	return rst
}
