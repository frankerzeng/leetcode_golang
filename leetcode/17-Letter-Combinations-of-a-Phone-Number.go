/*
17. Letter Combinations of a Phone Number
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

九宫格输入法

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package main

func main() {
	letterCombinations("234")
}
func letterCombinations(digits string) []string {
	var rst []string
	var keyMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	for k, _ := range digits {
		tmp := keyMap[digits[k:k+1]]
		if len(rst) == 0 {
			for kk, _ := range tmp {
				rst = append(rst, tmp[kk:kk+1])
			}
		} else {
			rstCopy := rst   // var copy rst
			rst = []string{} // set as empty
			for kkk, _ := range rstCopy {
				for kk, _ := range tmp {
					rst = append(rst, rstCopy[kkk]+tmp[kk:kk+1])
				}
			}
		}
	}

	return rst
}
