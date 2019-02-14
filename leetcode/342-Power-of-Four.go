/*
342. Power of Four
Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example 1:

Input: 16
Output: true
Example 2:

Input: 5
Output: false
Follow up: Could you solve it without loops/recursion?
*/
package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(-2147483648))
}

/*
二进制中
1 -> 1
4 -> 100
16-> 10000
64-> 1000000
....
....
1都在奇数位且只有一个1

num&1431655765==num 保证1都在奇数位
num&(num-1)==0 保证是只有一个1
综上就只有一个奇数位的1 如 100 是4
*/
func isPowerOfFour(num int) bool {
	return num != 0 && num&(num-1) == 0 && num&1431655765 == num
}
