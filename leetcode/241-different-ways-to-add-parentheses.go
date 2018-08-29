/*
241. Different Ways to Add Parentheses

Given a string of numbers and operators, return all possible results from computing all the different possible ways to group
numbers and operators. The valid operators are +, - and *.

Example 1:
	Input: "2-1-1"
	Output: [0, 2]
	Explanation:
	((2-1)-1) = 0
	(2-(1-1)) = 2

Example 2:
	Input: "2*3-4*5"
	Output: [-34, -14, -10, -10, 10]
	Explanation:
	(2*(3-(4*5))) = -34
	((2*3)-(4*5)) = -14
	((2*(3-4))*5) = -10
	(2*((3-4)*5)) = -10
	(((2*3)-4)*5) = 10
*/

package main

import (
	"fmt"
	"strconv"
)

// solution 4ms
func diffWaysToCompute(input string) []int {
	var oper []string // operator
	var nums []string // all numeric
	var flag = false

	for k := range input {
		if input[k:k+1] == "+" || input[k:k+1] == "-" || input[k:k+1] == "*" {
			flag = false
			oper = append(oper, input[k:k+1])
		} else {
			if flag {
				nums[len(nums)-1] = nums[len(nums)-1] + input[k:k+1]
			} else {
				nums = append(nums, input[k:k+1])
			}
			flag = true
		}
	}

	return dFun(nums, oper, 0, len(nums)-1)
}

// divide "nums" to two part，and the answer become the different way to combine with two part，and recursion this function
func dFun(nums []string, oper []string, start int, end int) []int {
	var rst []int
	if start == end {
		tmpInt, _ := strconv.Atoi(nums[start])
		rst = append(rst, tmpInt)
	} else {
		for i := start; i < end; i++ {
			tmp1 := dFun(nums, oper, start, i)
			tmp2 := dFun(nums, oper, i+1, end)

			// operate every numeric of two part
			for _, v := range tmp1 {
				for _, vv := range tmp2 {
					rst = append(rst, operateNums(v, vv, oper[i]))
				}
			}
		}
	}
	return rst
}

// operate two numeric
func operateNums(i int, j int, operate string) int {
	if operate == "+" {
		return i + j
	} else if operate == "-" {
		return i - j
	} else {
		return i * j
	}
}

func main() {
	input := "3*3+1*2"
	fmt.Println(diffWaysToCompute(input))
}
