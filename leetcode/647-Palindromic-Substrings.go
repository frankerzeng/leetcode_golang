/*
647. Palindromic Substrings

Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
Note:
The input string length won't exceed 1000.
*/
package main

import "fmt"

func main() {
	s := "sdd"
	fmt.Println(countSubstrings(s))
}
func countSubstrings(s string) int {
	rst := 0

	// whether ss is palindromic
	isPalindromic := func(ss string) bool {
		lenSS := len(ss)
		if lenSS == 1 {
			return true
		}
		mid := 0
		if lenSS%2 == 0 {
			mid = lenSS / 2
		} else {
			mid = (lenSS - 1) / 2
		}

		for i := 0; i < mid; i++ {
			if ss[i] != ss[lenSS-1-i] {
				return false
			}
		}
		return true
	}

	lenS := len(s)

	for j := 0; j < lenS; j++ {
		for k := j; k < lenS; k++ {
			if isPalindromic(s[j : k+1]) {
				rst++
			}
		}
	}

	return rst
}
