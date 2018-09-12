/*
86. Partition List

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.

Example:
	Input: head = 1->4->3->2->5->2, x = 3
	Output: 1->2->2->4->3->5
*/
package main

import "fmt"
import "../Common"

type ListNode = Common.ListNode

func main() {
	// 测试数据
	var node1 *ListNode = new(ListNode)
	var node2 *ListNode = new(ListNode)
	var node3 *ListNode = new(ListNode)
	node1.Val = 1
	node2.Val = 5
	node3.Val = 3
	node1.Next = node2
	node2.Next = node3

	fmt.Println(partition(node1, 2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var arrLess []int
	var arrGreat []int

	if head == nil {
		return head
	}
	var currentNods []*ListNode
	currentNods = append(currentNods, head)
	tmpNode := new(ListNode)
	tmpNode = head

	for true {
		if tmpNode.Val < x {
			arrLess = append(arrLess, tmpNode.Val)
		} else {
			arrGreat = append(arrGreat, tmpNode.Val)
		}
		if tmpNode.Next == nil {
			break
		}
		tmpNode = tmpNode.Next
	}
	returnNode := new(ListNode)
	returnNode = nil
	for i := len(arrGreat) - 1; i >= 0; i-- {
		tmpNode = new(ListNode)
		tmpNode.Val = arrGreat[i]
		if returnNode == nil {
			tmpNode.Next = nil
			returnNode = tmpNode
		} else {
			tmpNode.Next = returnNode
			returnNode = tmpNode
		}
	}

	for i := len(arrLess) - 1; i >= 0; i-- {
		tmpNode = new(ListNode)
		tmpNode.Val = arrLess[i]
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
