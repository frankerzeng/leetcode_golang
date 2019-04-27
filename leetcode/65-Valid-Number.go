/*
65. Valid Number
Validate if a given string can be interpreted as a decimal number.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false

Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before
implementing one. However, here is a list of characters that can be in a valid decimal number:

Numbers 0-9
Exponent - "e"
Positive/negative sign - "+"/"-"
Decimal point - "."
Of course, the context of these characters also matters in the input.

Update (2015-02-10):
The signature of the C++ function had been updated. If you still see your function signature accepts a const char *
	argument, please click the reload button to reset your code definition.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	trueCase := []string{
		"0",
		" 0.1 ",
		"2e10",
		" -90e3   ",
		" 6e-1",
		"53.5e93",
		"+.8",
		"3.",
	}
	falseCase := []string{
		"1 a",
		" --6 ",
		"-+3",
		"95a54e53",
		"-.",
		"6-1",
		"abc",
		" 1e",
		"e3",
		" 99e2.5 ",
	}
	for _, v := range trueCase {
		fmt.Println(v, isNumber(v))
	}
	fmt.Println("--------------------")
	for _, v := range falseCase {
		fmt.Println(v, isNumber(v))
	}
}
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	pointHas := false
	eHas := false
	numberHas := false
	numberAfterE := true
	for i := 0; i < len(s); i++ {
		if strings.Contains("0123456789", string(s[i])) {
			numberHas = true
			numberAfterE = true
		} else if s[i] == '.' {
			if eHas || pointHas { // 点前面不能有e
				return false
			}
			pointHas = true
		} else if s[i] == 'e' {
			if eHas || !numberHas { // e前面必须有数字
				return false
			}
			numberAfterE = false
			eHas = true
		} else if strings.Contains("-+", string(s[i])) {
			if i != 0 && s[i-1] != 'e' { // 加减号只能出现在第一位或者e后面
				return false
			}
		} else {
			return false
		}
	}
	return numberHas && numberAfterE
}
