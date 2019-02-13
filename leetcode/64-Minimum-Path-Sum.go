/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of
all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/
package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{7, 2}, {6, 6}, {8, 6}, {8, 7}, {5, 0}, {6, 0}}))
}
func minPathSum(grid [][]int) int {
	M = len(grid)
	N = len(grid[0])
	Grid = make([]int, M*N)
	return minSumFunc(grid, 1, 1)
}

var Grid []int
var M int
var N int

// 第m行 x n列
func minSumFunc(grid [][]int, m int, n int) int {
	var rst int
	if len(grid) == m {
		for k, v := range grid[m-1] {
			if k >= n-1 {
				rst += v
			}
		}
	} else if len(grid[0]) == n {
		for k, v := range grid {
			if k >= m-1 {
				rst += v[n-1]
			}
		}
	} else {
		tmp1 := Grid[m*N+n-1] // 避免重复计算
		if tmp1 == 0 {
			tmp1 = minSumFunc(grid, m+1, n)
			Grid[m*N+n-1] = tmp1
		}
		tmp2 := Grid[(m-1)*N+n]
		if tmp2 == 0 {
			tmp2 = minSumFunc(grid, m, n+1)
			Grid[(m-1)*N+n] = tmp2
		}
		if tmp1 > tmp2 {
			rst = grid[m-1][n-1] + tmp2
		} else {
			rst = grid[m-1][n-1] + tmp1
		}
	}

	return rst
}
