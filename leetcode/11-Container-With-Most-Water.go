/*
11. Container With Most Water
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are
drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a
container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section)
the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

// 暴力循环 faster then 5.00%
func maxArea1(height []int) int {
	var rst int
	lenHeight := len(height)
	for i := 0; i < lenHeight-1; i++ {
		for j := i + 1; j < lenHeight; j++ {
			tmp := (j - i) * int(math.Min(float64(height[j]), float64(height[i])))
			if rst < tmp {
				rst = tmp
			}
		}
	}

	return rst
}
func maxArea(height []int) int {
	var rst int
	var tmp int
	start := 0
	end := len(height) - 1
	for true {
		if start == end {
			break
		}
		if height[start] < height[end] {
			tmp = (end - start) * height[start]
		} else {
			tmp = (end - start) * height[end]
		}
		if rst < tmp {
			rst = tmp
		}
		if height[end] > height[start] {
			start++
		} else {
			end--
		}
	}

	return rst
}
