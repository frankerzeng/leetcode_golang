/**
598. Range Addition II
Given an m * n matrix M initialized with all 0's and several update operations.

Operations are represented by a 2D array, and each operation is represented by an array with two positive integers a and b,
which means M[i][j] should be added by one for all 0 <= i < a and 0 <= j < b.

You need to count and return the number of maximum integers in the matrix after performing all the operations.

Example 1:
Input:
m = 3, n = 3
operations = [[2,2],[3,3]]
Output: 4
Explanation:
Initially, M =
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]

After performing [2,2], M =
[[1, 1, 0],
 [1, 1, 0],
 [0, 0, 0]]

After performing [3,3], M =
[[2, 2, 1],
 [2, 2, 1],
 [1, 1, 1]]

So the maximum integer in M is 2, and there are four of it in M. So return 4.
*/
package main

import "fmt"

func main() {
	fmt.Println(maxCount(3, 3, [][]int{}))
}
func maxCount(m int, n int, ops [][]int) int {
	var mm = m // 结果是mm X nn
	var nn = n

	for _, v := range ops {
		mTmp := v[0]
		nTmp := v[1]
		if mTmp > m {
			mTmp = m
		}
		if nTmp > n {
			nTmp = n
		}
		if mm == 0 || mm > mTmp {
			mm = mTmp
		}
		if nn == 0 || nn > nTmp {
			nn = nTmp
		}
		fmt.Println(mm, nn)
	}

	return mm * nn
}
