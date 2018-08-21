/*
468. Validate IP Address
Write a function to check whether an input string is a valid IPv4 address or IPv6 address or neither.

IPv4 addresses are canonically represented in dot-decimal notation, which consists of four decimal numbers, each ranging
from 0 to 255, separated by dots ("."), e.g.,172.16.254.1;

Besides, leading zeros in the IPv4 is invalid. For example, the address 172.16.254.01 is invalid.

IPv6 addresses are represented as eight groups of four hexadecimal digits, each group representing 16 bits. The groups are
separated by colons (":"). For example, the address 2001:0db8:85a3:0000:0000:8a2e:0370:7334 is a valid one. Also, we could
omit some leading zeros among four hexadecimal digits and some low-case characters in the address to upper-case ones,
so 2001:db8:85a3:0:0:8A2E:0370:7334 is also a valid IPv6 address(Omit leading zeros and using upper cases).

However, we don't replace a consecutive group of zero value with a single empty group using two consecutive colons (::) to
pursue simplicity. For example, 2001:0db8:85a3::8A2E:0370:7334 is an invalid IPv6 address.

Besides, extra leading zeros in the IPv6 is also invalid. For example, the address 02001:0db8:85a3:0000:0000:8a2e:0370:7334 is
invalid.

Note: You may assume there is no extra space or special characters in the input string.

Example 1:
Input: "172.16.254.1"

Output: "IPv4"

Explanation: This is a valid IPv4 address, return "IPv4".
Example 2:
Input: "2001:0db8:85a3:0:0:8A2E:0370:7334"

Output: "IPv6"

Explanation: This is a valid IPv6 address, return "IPv6".
Example 3:
Input: "256.256.256.256"

Output: "Neither"

Explanation: This is neither a IPv4 address nor a IPv6 address.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("20EE:FGb8:85a3:0:0:8A2E:0370:7334"))
}

func validIPAddress(IP string) string {
	arr := strings.Split(IP, ".")
	ipv4 := "0123456789"       // ipv4 contains
	ipv6 := "0123456789ABCDEF" // ipv6 contains
	if len(arr) == 4 {
		for _, v := range arr {
			if v == "" || (v[0:1] == "0" && len(v) > 1) || len(v) > 3 {
				return "Neither"
			}
			for i := 0; i < len(v); i++ {
				if !strings.Contains(ipv4, v[0:1]) {
					return "Neither"
				}
			}
			if 3 == len(v) {
				tmpInt, _ := strconv.Atoi(v)
				if 255 < tmpInt || tmpInt == 0 { // max value is 255
					return "Neither"
				}
			}
		}
		return "IPv4"
	}

	IP = strings.ToUpper(IP)
	arr = strings.Split(IP, ":")
	if len(arr) == 8 {
		for _, v := range arr {
			if v == "" || len(v) > 4 {
				return "Neither"
			}
			for i := 0; i < len(v); i++ {
				if !strings.Contains(ipv6, v[i:i+1]) {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}

	return "Neither"
}
