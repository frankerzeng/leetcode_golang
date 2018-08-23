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
			// is pre is greater than current, we subtract with 10 or 100 or 1000 ...
			// then replace the end of string with 9 or 99 or 999 ...
			if pre > current {
				newN, _ := strconv.Atoi(NStr)
				tmp := 10
				for j := i; j < NLen-1; j++ {
					if pre > 1 {
						tmp *= 10
					}
				}
				NStr = strconv.Itoa(newN - tmp)
				tmpStr := ""
				for j := i; j < NLen; j++ {
					tmpStr += "9"
				}

				// that pre numeric is less then 1,subtraction 10^n will change the pre numeric that pre numeric
				if pre < 2 && len(tmpStr) > 1 {
					isEqu := NLenTemp == len(NStr)
					NLenTemp = len(NStr)
					if isEqu {
						NStr = NStr[0:i+1] + tmpStr[:len(tmpStr)-1]
						continue
					}
				}
				NStr = NStr[0:i] + tmpStr
			}
		}
	}
	n, _ := strconv.Atoi(NStr)
	return n
}

func main() {
	fmt.Println(monotoneIncreasingDigits(332))
}
