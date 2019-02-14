/*
52. N-Queens II
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

Example:

Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(totalNQueens(4))
}

// 表格长度
var gridLen int

// 对应的没一列在第几行的位置
var gridCurr []int

// nxn 的表格，1表示放入皇后
var gridValidate [][]int

func totalNQueens(n int) int {
	gridLen = n
	gridValidate = make([][]int, n)
	for k := range gridValidate {
		gridValidate[k] = make([]int, n)
		gridCurr = append(gridCurr, -1)
	}
	var rst = 0
	setFunc(0, 0, &rst)
	return rst
}

// m行n列 是否可以放
func setFunc(m int, n int, rst *int) {
	if m == gridLen && n == 0 { // 遍历结束
		return
	}
	if m == gridLen {
		gridValidate[gridCurr[n-1]][n-1] = 0
		setFunc(gridCurr[n-1]+1, n-1, rst)
		return
	}
	if n == gridLen {
		gridValidate[gridCurr[n-1]][n-1] = 0
		setFunc(gridCurr[n-1]+1, n-1, rst)
		return
	}
	tmp := isValidate(m, n)
	gridValidate[m][n] = 0

	if tmp {
		gridCurr[n] = m
		gridValidate[m][n] = 1
		if n == gridLen-1 {
			*rst = *rst + 1
		}
		setFunc(0, n+1, rst)
	} else {
		setFunc(m+1, n, rst)
	}
}

// m行n列 是否可以放
func isValidate(m int, n int) bool {
	for i := 0; i < len(gridValidate); i++ { // i行
		for j := 0; j < len(gridValidate); j++ { // j列
			if gridValidate[i][j] == 1 {
				if i == m || j == n || math.Abs(float64(i-m)) == math.Abs(float64(j-n)) {
					return false
				}
			}
		}
	}
	return true
}
