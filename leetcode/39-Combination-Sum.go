/*
39. Combination Sum
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

*/
package main

import "fmt"

func main() {

	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
func combinationSum(candidates []int, target int) [][]int {
	var rst [][]int
	for i := 0; i < len(candidates); i++ {
		if target-candidates[i] > 0 {
			for _, rst_ := range combinationSum(candidates[i:], target-candidates[i]) {
				tmp := make([]int, len(rst_)+1)
				copy(tmp, rst_)
				tmp[len(rst_)] = candidates[i]
				rst = append(rst, tmp)
			}
		} else if target-candidates[i] == 0 {
			rst = append(rst, []int{candidates[i]})
		}
	}

	return rst
}
