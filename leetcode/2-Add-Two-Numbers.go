/*
2. Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order
and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package main

import (
	"../Common"
	"fmt"
	"strconv"
)

type ListNode = Common.ListNode

func main() {
	l1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp2 := ""
	flag := false // 是否进位
	for true {
		if l1 == nil && l2 == nil {
			break
		}
		i1 := 0
		i2 := 0
		if l1 != nil {
			i1 = l1.Val
		}
		if l2 != nil {
			i2 = l2.Val
		}
		i12 := i1 + i2
		if flag {
			i12++
		}
		if i12 > 9 {
			flag = true
		} else {
			flag = false
		}

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
		tmp2 = strconv.Itoa(i12%10) + tmp2
	}
	if flag {
		tmp2 = "1" + tmp2
	}

	rst := &ListNode{}
	for i := 0; i < len(tmp2); i++ {
		strTmp, _ := strconv.Atoi(tmp2[i : i+1])
		if i == 0 {
			rst = &ListNode{
				Val: strTmp,
			}
		} else {
			rst = &ListNode{
				Val:  strTmp,
				Next: rst,
			}
		}
	}

	tmp1 := rst
	for true {
		if tmp1.Next != nil {
			tmp1 = tmp1.Next
		} else {
			break
		}
	}
	return rst
}
