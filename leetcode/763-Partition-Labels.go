/*
763. Partition Labels
A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

Example 1:
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
Note:

S will have length in range [1, 500].
S will consist of lowercase letters ('a' to 'z') only.
*/
package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(S string) []int {
	var rst []int

	oneFunc := func(S string) int { // 返回第一个长度
		rst := 1
		for i := 0; i < len(S); i++ {
			if i > rst-1 {
				break
			}
			for j := len(S) - 1; j >= i+1; j-- {
				if S[j] == S[i] && rst < j+1 {
					rst = j + 1
					break
				}
			}
		}

		return rst
	}

	for len(S) > 0 {
		oneFuncTmp := oneFunc(S)
		rst = append(rst, oneFuncTmp)
		S = S[oneFuncTmp:]
	}

	return rst
}
