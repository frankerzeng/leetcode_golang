/*
143. Reorder List
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
*/
package main

import (
	"fmt"
)
import (
	common "../Common"
)

type ListNode = common.ListNode

func main() {
	head := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: nil,
				},
			},
		},
	}
	fmt.Println(head)
	head = nil
	reorderList(head)
	common.PrintList(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	move(head, 1)
}

// 进行第几次移动，把最后一个节点移动到当前位置，index为当前待移动位置
func move(head *ListNode, index int) {
	if head == nil {
		return
	}
	var curr int = 1
	var currNode = head
	var headNode = head          // 移动到此节点后面
	var lastSecondNode *ListNode // 待移动节点的前一个节点
	var lastNode *ListNode       // 待移动节点
	for true {
		if currNode.Next != nil && currNode.Next.Next == nil {
			lastSecondNode = currNode
		}
		if currNode.Next == nil {
			lastNode = currNode
			break
		}
		if curr == index {
			headNode = currNode
		}
		currNode = currNode.Next
		curr++
	}
	if curr-2 < index {
		return
	}

	lastSecondNode.Next = nil
	lastNode.Next = headNode.Next
	headNode.Next = lastNode

	move(head, index+2)
}
