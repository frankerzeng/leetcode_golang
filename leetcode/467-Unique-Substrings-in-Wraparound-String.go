/**
467. Unique Substrings in Wraparound String
Consider the string s to be the infinite wraparound string of "abcdefghijklmnopqrstuvwxyz", so s will look like this:
"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".

Now we have another string p. Your job is to find out how many unique non-empty substrings of p are present in s. In
particular, your input is the string p and you need to output the number of different non-empty substrings of p in the
string s.

Note: p consists of only lowercase English letters and the size of p might be over 10000.

Example 1:
Input: "a"
Output: 1
Explanation: Only the substring "a" of string "a" is in the string s.

Example 2:
Input: "cac"
Output: 2
Explanation: There are two substrings "a", "c" of string "cac" in the string s.

Example 3:
Input: "zab"
Output: 6
Explanation: There are six substrings "z", "a", "b", "za", "ab", "zab" of string "zab" in the string s.

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findSubstringInWraproundString("cac"))
}

/*
思路：
找出连续的最大的substring
abce => abc 和 e

*/
func findSubstringInWraproundString(p string) int {
	var rst int
	if len(p) == 0 {
		return rst
	}

	rstSlice := [26]int{} // a开头的，b开头的

	for i := 0; i < len(p); {
		index := i + 1
		for index < len(p) && (p[index] == p[index-1]+1 || p[index] == p[index-1]-25) {
			index++
		}
		for j := i; j < index; j++ {
			id := int(p[j]) - int('a')
			rstSlice[id] = int(math.Max(float64(rstSlice[id]), float64(index-j)))
		}
		i = index
	}

	for _, v := range rstSlice {
		rst += v
	}

	return rst
}
