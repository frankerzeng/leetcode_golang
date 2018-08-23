/*
738. Monotone Increasing Digits
Given a non-negative integer N, find the largest number that is less than or equal to N with monotone increasing digits.

(Recall that an integer has monotone increasing digits if and only if each pair of adjacent digits x and y satisfy x <= y.)

Example 1:
Input: N = 10
Output: 9
Example 2:
Input: N = 1234
Output: 1234
Example 3:
Input: N = 332
Output: 299
Note: N is an integer in the range [0, 10^9].
*/

package main

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits(N int) int {
	NStr := strconv.Itoa(N)
	NLen := len(NStr)
	NLenTemp := NLen

	for i := NLen - 1; i >= 0; i-- {
		if i-1 >= 0 {
			current, _ := strconv.Atoi(NStr[i : i+1]) // current numeric
			pre, _ := strconv.Atoi(NStr[i-1 : i])     // pre numeric
			if pre > current {
				newN, _ := strconv.Atoi(NStr)
				tmp := 10
				for j := i; j < NLen-1; j++ {
					if pre > 1 {
						tmp *= 10
					}
				}
				newN = newN - tmp
				NStr = strconv.Itoa(newN)
				tmpStr := ""
				for j := i; j < NLen; j++ {
					tmpStr += "9"
				}

				if pre < 2 && len(tmpStr) > 1 {
					tmpStr = tmpStr[:len(tmpStr)-1]
					if NLenTemp == len(NStr) {
						NStr = NStr[0:i+1] + tmpStr
					} else {
						NLenTemp = len(NStr)
						NStr = NStr[0:i] + tmpStr
					}
				} else {
					NStr = NStr[0:i] + tmpStr
				}
			}
		}
	}
	n, _ := strconv.Atoi(NStr)
	return n
}

func main() {
	fmt.Println(monotoneIncreasingDigits(332))
}
