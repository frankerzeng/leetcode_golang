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
	fmt.Println(isInterleave("aabcc",
		"dbbc a",
		"aadbbcbca c"))

}
func isInterleave(s1 string, s2 string, s3 string) bool {
	s3Index := 0 // the comparing character index of s3
	s2Index := 0 // the comparing character index of s2
	s1Index := 0 // the comparing character index of s1
	s3Len := len(s3)
	s2Len := len(s2)
	s1Len := len(s1)
	for true {
		if s3Index == s3Len-1 {
			return true
		}
		fmt.Println(s1Index, s2Index, s3Index)

		if s1Index < s1Len && s3[s3Index:s3Index+1] == s1[s1Index:s1Index+1] {
			s3Index++
			s1Index++
			continue
		}
		if s2Index < s2Len && s3[s3Index:s3Index+1] == s2[s2Index:s2Index+1] {
			s3Index++
			s2Index++
			continue
		}

		if s3Index > 1 && s1Index > 1 {
			if s3[s3Index-1:s3Index] == s2[s1Index-1:s1Index] && s3[s3Index-1:s3Index] == s2[s2Index-1:s2Index] {
				s2Index++
				s1Index--
			} else {

			}
		}
		fmt.Println("---------")

	}
	return true
}
