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
	fmt.Println(findKthNumber(10000, 10000))
}

// 1 10 100 11 12 13 14 15 16 17
// 1
// 10  11 12 13 --- 19
// 100 101 ---109 ---199
func findKthNumber(n int, k int) int {
	fmt.Println("109" < "110")
	k--
	var rstInt int
	for i := 1; i < 10; i++ {
		tmp := i
		fmt.Println(tmp, 1)
		if tmp > n {
			continue
		}
		if k == 0 {
			return tmp
		}
		k--
		for j := 0; j < 10; j++ {
			tmp1 := tmp*10 + j
			fmt.Println(tmp1, 2)
			if tmp1 > n {
				break
			}
			if k == 0 {
				return tmp1
			}
			k--
			for h := 0; h < 10; h++ {
				tmp2 := tmp1*10 + h
				fmt.Println(tmp2, 3)
				if tmp2 > n {
					break
				}
				if k == 0 {
					return tmp2
				}
				k--
			}
		}
	}

	return rstInt
}

func rool() int {

	return -1
}
