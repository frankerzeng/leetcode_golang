/*
440. K-th Smallest in Lexicographical Order
Given integers n and k, find the lexicographically k-th smallest integer in the range from 1 to n.

Note: 1 ≤ k ≤ n ≤ 10^9

Example:

Input:
n: 13   k: 2

Output:
10

Explanation:
The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthNumber(13, 4))
}

// 1 10 100 11 12 13 14 15 16 17
// 1
// 10  11 12 13 --- 19
// 100 101 ---109 ---199
func findKthNumber(n int, k int) int {
	curr := 1
	k--

	for k > 0 {
		steps := calSteps(n, curr, curr+1)
		if steps <= k {
			curr++
			k -= steps
		} else {
			curr *= 10
			k--
		}
	}
	return curr
}

func calSteps(n int, n1 int, n2 int) int {
	fmt.Println(n, n1, n2)
	steps := 0
	for n1 <= n {
		if (n + 1) < n2 {
			steps += n + 1 - n1
		} else {
			steps += n2 - n1
		}
		n1 *= 10
		n2 *= 10
	}
	return steps
}
