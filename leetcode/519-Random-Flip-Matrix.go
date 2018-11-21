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
	obj := Constructor(2, 2)
	fmt.Println("RESET==================")

	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	return

	for _, v := range []string{"flip", "flip",
		"reset",
		"flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip",
		"reset",
		"reset",
		"flip", "flip", "reset", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip",
		"reset",
		"flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip",
		"reset",
		"flip", "flip", "flip", "flip", "flip", "flip", "flip", "reset", "flip",
		"reset",
		"flip", "flip", "flip", "flip", "flip", "flip",
		"reset",
		"flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip", "flip",
		"reset",
		"flip", "flip", "flip", "flip",
	} {
		if v == "flip" {
			fmt.Println(obj.Flip())
		} else if v == "reset" {
			fmt.Println("RESET==================")
			obj.Reset()
			fmt.Println(obj)
		}
	}
}

type Solution struct {
	MaxLength     int
	ValRow        int
	canNotRandMap []int
}

func Constructor(n_rows int, n_cols int) Solution {
	n_rows, n_cols = n_cols, n_rows
	return Solution{MaxLength: n_rows*n_cols - 1, canNotRandMap: []int{}, ValRow: n_rows}
}

func (this *Solution) Flip() []int {
	row, col, flag, randInt := 0, 0, false, 0
	for true {
		randInt = rand.Intn(this.MaxLength + 1)
		flag = false
		for _, v := range this.canNotRandMap {
			if v == randInt {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		this.canNotRandMap = append(this.canNotRandMap, randInt)
		break
	}
	col = randInt % this.ValRow
	row = (randInt - col) / this.ValRow
	return []int{row, col}

}

func (this *Solution) Reset() {
	this.canNotRandMap = []int{}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
