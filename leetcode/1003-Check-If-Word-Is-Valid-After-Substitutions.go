/*
1003. Check If Word Is Valid After Substitutions
We are given that the string "abc" is valid.

From any valid string V, we may split V into two pieces X and Y such that X + Y (X concatenated with Y) is equal to V.  (X or Y may be empty.)  Then, X + "abc" + Y is also valid.

If for example S = "abc", then examples of valid strings are: "abc", "aabcbc", "abcabc", "abcabcababcc".  Examples of invalid strings are: "abccba", "ab", "cababc", "bac".

Return true if and only if the given string S is valid.



Example 1:

Input: "aabcbc"
Output: true
Explanation:
We start with the valid string "abc".
Then we can insert another "abc" between "a" and "bc", resulting in "a" + "abc" + "bc" which is "aabcbc".
Example 2:

Input: "abcabcababcc"
Output: true
Explanation:
"abcabcabc" is valid after consecutive insertings of "abc".
Then we can insert "abc" before the last letter, resulting in "abcabcab" + "abc" + "c" which is "abcabcababcc".
Example 3:

Input: "abccba"
Output: false
Example 4:

Input: "cababc"
Output: false


Note:

1 <= S.length <= 20000
S[i] is 'a', 'b', or 'c'
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("aabcbc") == true)
	fmt.Println(isValid("abcabcababcc") == true)
	fmt.Println(isValid("abccba") == false)
	fmt.Println(isValid("cababc") == false)
}

// 把所有连续的abc剔除掉如果剩下的是"abc"或者""则返回true
func isValid(S string) bool {
	Stmp := S
	for true {
		Stmp = S
		S = strings.Replace(S, "abc", "", -1)
		if S == Stmp {
			break
		}
	}

	return S == ""
}
