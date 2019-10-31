/*
686. Repeated String Match
Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("aa", "a"))
}

func repeatedStringMatch(A string, B string) int {
	rst, originalA := -1, A
	minI := len(B)/len(A) - 1 // 最小循环次数
	if minI < 0 {
		minI = 0
	}
	maxI := len(B)/len(A) + 1 // 最大循环次数

	for i := 1; i <= minI; i++ {
		A = A + originalA
	}

	for i := minI; i <= maxI; i++ {
		if strings.Contains(A, B) {
			return i + 1
		}
		A = A + originalA
	}

	return rst
}
