/*
97. Interleaving String

Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

Example 1:
	Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
	Output: true

Example 2:
	Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
	Output: false
*/
package main

import "fmt"

func main() {
	//fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))

}
func isInterleave(s1 string, s2 string, s3 string) bool {
	//s3Index  the comparing character index of s3
	//s2Index  the comparing character index of s2
	//s1Index  the comparing character index of s1
	if len(s1)+len(s1) != len(s3) {
		return false
	}
	return dpFunc(s1, s2, s3, 0, 0, 0)
}
func dpFunc(s1 string, s2 string, s3 string, s1Index int, s2Index int, s3Index int) bool {
	if s3Index == len(s3) {
		return true
	}
	if s1Index < len(s1) {
		if s1[s1Index:s1Index+1] == s3[s3Index:s3Index+1] && dpFunc(s1, s2, s3, s1Index+1, s2Index, s3Index+1) {
			return true
		}
	}
	if s2Index < len(s2) {
		if s2[s2Index:s2Index+1] == s3[s3Index:s3Index+1] && dpFunc(s1, s2, s3, s1Index, s2Index+1, s3Index+1) {
			return true
		}
	}
	return false
}
