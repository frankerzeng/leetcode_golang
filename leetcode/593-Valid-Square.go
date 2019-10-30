/*
593. Valid Square
Given the coordinates of four points in 2D space, return whether the four points could construct a square.

The coordinate (x,y) of a point is represented by an integer array with two integers.

Example:

Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
Output: True


Note:

All the input integers are in the range [-10000, 10000].
A valid square has four equal sides with positive length and four equal angles (90-degree angles).
Input points have no order.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	// [1,0],[-1,0],[0,1],[0,-1]
	// [0,0] [5,0] [5,4] [0,4]
	// [0,0][1,1][1,0][0,1]
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 任意两个点之间的距离应该一样
	disEqu := func(p1 []int, p2 []int, p3 []int, p4 []int) bool {
		if (p1[0] == p2[0] && p1[1] == p2[1]) || (p3[0] == p4[0] && p3[1] == p4[1]) {
			return false
		}
		if math.Pow(float64(p1[0])-float64(p2[0]), float64(2))+math.Pow(float64(p1[1])-float64(p2[1]), float64(2)) ==
			math.Pow(float64(p3[0])-float64(p4[0]), float64(2))+math.Pow(float64(p3[1])-float64(p4[1]), float64(2)) {
			return true
		}
		return false
	}

	if disEqu(p1, p2, p3, p4) && disEqu(p1, p4, p2, p3) && disEqu(p1, p3, p2, p4) && // 任何两点距离相等
		(disEqu(p1, p2, p1, p3) || disEqu(p1, p4, p1, p3) || disEqu(p1, p4, p1, p2)) { // 相邻两条边相等
		return true
	}

	return false
}
