/*
83. Remove Duplicates from Sorted List

Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
	Input: 1->1->2
	Output: 1->2

Example 2:
	Input: 1->1->2->3->3
	Output: 1->2->3
*/

package main

import (
	"fmt"
)
import "../Common"

type ListNode = Common.ListNode

func main() {
	// 测试数据
	var node1 *ListNode = new(ListNode)

	node1 = nil

	fmt.Println(deleteDuplicates(node1))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var currentNods []*ListNode
	currentNods = append(currentNods, head)

	tmpNode := new(ListNode)
	tmpNode = head

	var intMap []int
	flag := false

	for true {
		flag = false
		for _, vv := range intMap {
			if vv == tmpNode.Val {
				flag = true
				break
			}
		}
		if !flag {
			intMap = append(intMap, tmpNode.Val)
		}
		if tmpNode.Next == nil {
			break
		}
		tmpNode = tmpNode.Next
	}

	returnNode := new(ListNode)
	returnNode = nil
	for i := len(intMap) - 1; i >= 0; i-- {
		tmpNode = new(ListNode)
		tmpNode.Val = intMap[i]
		if returnNode == nil {
			tmpNode.Next = nil
			returnNode = tmpNode
		} else {
			tmpNode.Next = returnNode
			returnNode = tmpNode
		}
	}

	return returnNode
}
