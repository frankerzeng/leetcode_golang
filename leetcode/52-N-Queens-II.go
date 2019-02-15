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

// DFS 算法
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

// 放在m行n列
func setFunc(m int, n int, rst *int) {
	if m == gridLen && n == 0 { // 遍历到(gridlen,0)这表明遍历结束
		return
	}
	if m == gridLen {
		gridValidate[gridCurr[n-1]][n-1] = 0
		setFunc(gridCurr[n-1]+1, n-1, rst) // 超出范围，上一列走下一步
		return
	}
	if n == gridLen {
		gridValidate[gridCurr[n-1]][n-1] = 0
		setFunc(gridCurr[n-1]+1, n-1, rst) // 超出范围，上一列走下一步
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
		setFunc(0, n+1, rst) // 本列，走下一步，肯定在下一列
	} else {
		setFunc(m+1, n, rst) // 超出范围，本列走下一步
	}
}

// m行n列 是否可以放
func isValidate(m int, n int) bool {
	// current columns and current rows
	for i := 0; i < gridLen; i++ {
		if gridValidate[m][i] == 1 || gridValidate[i][n] == 1 {
			return false
		}
	}

	// diagonal
	for i := 0; i < gridLen; i++ {
		if m+i < gridLen && n+i < gridLen {
			if gridValidate[m+i][n+i] == 1 {
				return false
			}
		}
		if m+i < gridLen && n-i > -1 {
			if gridValidate[m+i][n-i] == 1 {
				return false
			}
		}
		if m-i > -1 && n+i < gridLen {
			if gridValidate[m-i][n+i] == 1 {
				return false
			}
		}
		if m-i > -1 && n-i > -1 {
			if gridValidate[m-i][n-i] == 1 {
				return false
			}
		}
	}

	return true
}
