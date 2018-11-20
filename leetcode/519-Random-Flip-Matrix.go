/*
519. Random Flip Matrix
You are given the number of rows n_rows and number of columns n_cols of a 2D binary matrix where all values are initially 0.
Write a function flip which chooses a 0 value uniformly at random, changes it to 1, and then returns the position
[row.id, col.id] of that value. Also, write a function reset which sets all values back to 0. Try to minimize the number of
calls to system's Math.random() and optimize the time and space complexity.

Note:
	1 <= n_rows, n_cols <= 10000
	0 <= row.id < n_rows and 0 <= col.id < n_cols
	flip will not be called when the matrix has no 0 values left.
	the total number of calls to flip and reset will not exceed 1000.

Example 1:
Input:
	["Solution","flip","flip","flip","flip"]
	[[2,3],[],[],[],[]]
Output:
	[null,[0,1],[1,2],[1,0],[1,1]]

Example 2:
Input:
	["Solution","flip","flip","reset","flip"]
	[[1,2],[],[],[],[]]
Output:
	[null,[0,0],[0,1],null,[0,0]]

Explanation of Input Syntax:

The input is two lists: the subroutines called and their arguments. Solution's constructor has two arguments, n_rows
and n_cols. flip and reset have no arguments. Arguments are always wrapped with a list, even if there aren't any.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor(3, 2)
	fmt.Println(obj)
	param1 := obj.Flip()
	fmt.Println(param1)
	param1 = obj.Flip()
	fmt.Println(param1)
	param1 = obj.Flip()
	fmt.Println(param1)
	param1 = obj.Flip()
	fmt.Println(param1)
	param1 = obj.Flip()
	fmt.Println(param1)
	param1 = obj.Flip()
	fmt.Println(param1)
	param1 = obj.Flip()
	fmt.Println(param1)
	obj.Reset()
	fmt.Println(obj)
}

type Solution struct {
	Val    [][]int
	lenRow int
	lenCol int
	OneRow int // current 1 index
	OneCol int // current 1 index
}

func Constructor(n_rows int, n_cols int) Solution {
	m1 := make([][]int, n_rows, n_rows)
	for i := 0; i < n_rows; i++ {
		m1[i] = make([]int, n_cols, n_cols)
	}
	return Solution{Val: m1, lenRow: n_rows, lenCol: n_cols, OneRow: -1, OneCol: -1}
}

func (this *Solution) Flip() []int {
	var rst []int
	if this.OneCol == -1 && this.OneRow == -1 {
		rst = []int{0, 0}
		this.OneCol++
		this.OneRow++
	} else {
		if this.OneCol < this.lenCol-1 {
			this.OneCol++
		} else if this.OneRow < this.lenRow-1 {
			this.OneRow++
			this.OneCol = 0
		}
		rst = []int{this.OneRow, this.OneCol}
	}
	contain
	this.Val[this.OneRow][this.OneCol] = 1
	return rst

}

func (this *Solution) Reset() {
	*this = Constructor(2, 2)
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
