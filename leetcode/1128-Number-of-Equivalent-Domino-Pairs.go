/*
1128. Number of Equivalent Domino Pairs
Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that is, one domino can be rotated to be equal to another domino.

Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].



Example 1:

Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
Output: 1
*/
package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	rst := 0
	lenDominoes := len(dominoes)
	for k, v := range dominoes {
		for i := k + 1; i < lenDominoes; i++ {
			if (v[0] == dominoes[i][0] && v[1] == dominoes[i][1]) || (v[1] == dominoes[i][0] && v[0] == dominoes[i][1]) {
				rst++
			}
		}
	}
	return rst
}
func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}
