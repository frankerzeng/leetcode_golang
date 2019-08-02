/**
1037. Valid Boomerang
A boomerang is a set of 3 points that are all distinct and not in a straight line.

Given a list of three points in the plane, return whether these points are a boomerang.



Example 1:

Input: [[1,1],[2,3],[3,2]]
Output: true
Example 2:

Input: [[1,1],[2,2],[3,3]]
Output: false


Note:

points.length == 3
points[i].length == 2
0 <= points[i][j] <= 100
*/
package main

import "fmt"

func main() {
	// fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
	// fmt.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// fmt.Println(isBoomerang([][]int{{0, 0}, {0, 2}, {2, 1}}))
	fmt.Println(isBoomerang([][]int{{0, 0}, {2, 1}, {2, 1}}) == false)
	fmt.Println(isBoomerang([][]int{{0, 1}, {0, 1}, {2, 1}}) == false)
	fmt.Println(isBoomerang([][]int{{0, 1}, {2, 0}, {1, 1}}) == true)
	fmt.Println(isBoomerang([][]int{{1, 0}, {0, 0}, {2, 0}}) == false)
}

func isBoomerang(points [][]int) bool {
	if points[1][0]-points[0][0] != 0 &&
		float64(points[2][1]) == (float64(points[1][1]-points[0][1])/float64(points[1][0]-points[0][0]))*float64(points[2][0]-points[0][0])+float64(points[0][1]) { // 在一条直线上
		return false
	} else if points[1][0] == points[0][0] && points[2][0] == points[0][0] { // 在一条垂直线上
		return false
	} else if points[1][0] == points[0][0] && points[1][1] == points[0][1] { // 1,2点重叠了
		return false
	}
	return true
}
