/*
387. First Unique Character in a String
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/
package main

import (
	"fmt"
)

func main() {
	// fmt.Println(firstUniqChar("leetcode") == 0)
	fmt.Println(firstUniqChar("loveleetcode") == 2)
}
func firstUniqChar(s string) int {
	var rst = -1
	var rstMap = make(map[int32]int, 26)
	for k, v := range s {
		if _, ok := rstMap[v]; ok {
			rstMap[v] = -1
		} else {
			rstMap[v] = k
		}
	}
	for _, v := range rstMap {
		if rst == -1 || (v > -1 && v < rst) {
			rst = v
		}
	}
	return rst
}
