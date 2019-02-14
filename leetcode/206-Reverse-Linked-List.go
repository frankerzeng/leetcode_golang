/*
206. Reverse Linked List
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

package main

import (
	"fmt"
	"runtime"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	memConsumed := func() uint64 {
		runtime.GC() // GC，排除对象影响
		var memStat runtime.MemStats
		runtime.ReadMemStats(&memStat)
		return memStat.Sys
	}
	before := memConsumed()
	var node1 *ListNode = new(ListNode)
	var node2 *ListNode = new(ListNode)
	var node3 *ListNode = new(ListNode)
	var node4 *ListNode = new(ListNode)
	var node5 *ListNode = new(ListNode)
	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = nil
	fmt.Println(reverseList(node1))
	after := memConsumed()
	const goroutineNum = 1e4
	fmt.Printf("%.30f KB\n", float64(after-before))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode { // Memory Limit Exceeded ?????
	var prev *ListNode
	var curr = head
	for curr != nil {
		var nextTemp = curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode { // Memory Limit Exceeded ?????
	reverseNode(head)
	head.Next = nil
	return firstNode
}

var firstNode *ListNode

func reverseNode(node *ListNode) {
	if node.Next.Next != nil {
		reverseNode(node.Next)
	} else {
		firstNode = node.Next
		node.Next.Next = node
	}
}
