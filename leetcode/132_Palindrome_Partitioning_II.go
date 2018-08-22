/*
132. Palindrome Partitioning II
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

Example:

Input: "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/
package main

import "fmt"

func main() {
	fmt.Println(minCut("cabababcbc"))
}
func minCut(s string) int {
	m := 0
	palindromeMap := make(map[int]map[int]bool) // weather i to j is a palindrome
	minChar := make(map[int]int)                // mini cut for 0 to i string

	for i := 0; i < len(s); i++ {
		m = i
		for j := 0; j <= i; j++ {
			if s[j:j+1] == s[i:i+1] {
				_, v := palindromeMap[j+1][i-1]
				if v == true || j+1 > i-1 {
					palindromeMapChild := make(map[int]bool)
					palindromeMapChild[i] = true
					palindromeMap[j] = palindromeMapChild
					if j == 0 {
						m = 0
					} else {
						if m > (minChar[j-1] + 1) {
							m = minChar[j-1] + 1
						}
					}
				}
			}
		}
		minChar[i] = m
	}
	return minChar[len(s)-1]
}
