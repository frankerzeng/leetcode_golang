/*
537. Complex Number Multiplication
Given two strings representing two complex numbers.

You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

Example 1:
Input: "1+1i", "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
Example 2:
Input: "1+-1i", "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
Note:

The input strings will not have extra blank.
The input strings will be given in the form of a+bi, where the integer a and b will both belong to the range of [-100, 100].
And the output should be also in this form.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}

/**
分解：
(a+bi)*(c+di) = ac- bd + (bc+ad)i
*/
func complexNumberMultiply(a string, b string) string {
	a1, a2 := getAB(a)
	b1, b2 := getAB(b)
	return strconv.Itoa(a1*b1-a2*b2) + "+" + strconv.Itoa(a2*b1+a1*b2) + "i"
}

func getAB(A string) (a, b int) {
	indexPlus := strings.Index(A, "+")
	tmpA := A[0:indexPlus]
	tmpB := A[indexPlus+1 : len(A)-1]
	a, _ = strconv.Atoi(tmpA)
	b, _ = strconv.Atoi(tmpB)
	return
}
