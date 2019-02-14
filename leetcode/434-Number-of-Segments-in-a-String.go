/*
434. Number of Segments in a String
Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.

Please note that the string does not contain any non-printable characters.

Example:

Input: "Hello, my name is John"
Output: 5
*/
package main

import "fmt"

func main() {
	fmt.Println(countSegments(" ds  d fs"))

}

// 自己非空，前一个元素为空 则+1
func countSegments(s string) int {
	var rst int

	for k, v := range s {
		if k == 0 {
			if v != ' ' {
				rst++
			}
		} else {
			if s[k-1] == ' ' && v != ' ' {
				rst++
			}
		}
	}

	return rst
}
