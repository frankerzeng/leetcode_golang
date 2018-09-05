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
	node1 = nil

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

	fmt.Println(12)
	return head
}
