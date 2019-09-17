/**
1023. Camelcase Matching
A query word matches a given pattern if we can insert lowercase letters to the pattern word so that it equals the query.
(We may insert each character at any position, and may insert 0 characters.)

Given a list of queries, and a pattern, return an answer list of booleans, where answer[i] is true if and only if queries[i]
matches the pattern.

Example 1:
Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
Output: [true,false,true,true,false]
Explanation:
"FooBar" can be generated like this "F" + "oo" + "B" + "ar".
"FootBall" can be generated like this "F" + "oot" + "B" + "all".
"FrameBuffer" can be generated like this "F" + "rame" + "B" + "uffer".

Example 2:
Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
Output: [true,false,true,false,false]
Explanation:
"FooBar" can be generated like this "Fo" + "o" + "Ba" + "r".
"FootBall" can be generated like this "Fo" + "ot" + "Ba" + "ll".

Example 3:
Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
Output: [false,true,false,false,false]
Explanation:
"FooBarTest" can be generated like this "Fo" + "o" + "Ba" + "r" + "T" + "est".

*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"))
}
func camelMatch(queries []string, pattern string) []bool {
	var rst []bool
	var lenQuery int
	var lenPattern int
	var indexPattern int

	// 判断pattern是否可以组成query，全部小写字符
	for k := range queries {
		rst = append(rst, true)
		lenQuery = len(queries[k])
		lenPattern = len(pattern)
		if lenPattern > lenQuery {
			rst[k] = false
			continue
		}

		indexPattern = 0
		for i := 0; i < lenQuery; i++ {
			if indexPattern < lenPattern && queries[k][i] == pattern[indexPattern] {
				indexPattern++
			} else if queries[k][i] < 91 {
				rst[k] = false
				break
			}
		}
		if rst[k] == true && indexPattern != lenPattern {
			rst[k] = false
		}
	}

	return rst
}
