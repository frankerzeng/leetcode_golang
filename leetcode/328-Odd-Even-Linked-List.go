/*
328. Odd Even Linked List
Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:

Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
Example 2:

Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
Note:

The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
*/
package main

import (
	"fmt"
	"github.com/frankerzeng/golang_demo/Common"
)

type ListNode = Common.ListNode

func main() {
	head := &ListNode{}
	curr := head
	for _, v := range []int{1, 2, 3, 4} {
		tmp := &ListNode{
			Val: v,
		}
		curr.Next = tmp
		curr = curr.Next
	}
	head = head.Next
	prin(oddEvenList(head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	currentNode := head
	isOdd := true       // 是否是奇数位
	lastOddNode := head // 奇数位的最后一个节点
	if lastOddNode == nil {
		return head
	}
	firstEvenNode := head.Next // 偶数位的第一个节点
	if firstEvenNode == nil {
		return head
	}

	first := true
	for currentNode != nil { // 遍历链表
		if isOdd {
			if currentNode.Next != nil && currentNode.Next.Next != nil {
				tmpVal := &ListNode{ // 取出的奇数位的值
					Val:  currentNode.Next.Next.Val,
					Next: firstEvenNode,
				}

				currentNode.Next.Next = currentNode.Next.Next.Next
				lastOddNode.Next = tmpVal      // 找到一个奇数位就把奇数位插入到最后一个奇数位的位置
				lastOddNode = lastOddNode.Next // 同时跟新最后一个奇数位的位置

				if !first { // 非第一次会插入到current node 的前面，打乱顺序，所以取反
					isOdd = !isOdd
				}
				first = false
			} else {
				break
			}
		}
		isOdd = !isOdd
		currentNode = currentNode.Next
	}
	return head
}

func prin(head *ListNode) {
	currentNode := head
	var intSlice []int
	for currentNode != nil {
		intSlice = append(intSlice, currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println(intSlice)
}
