/*
235. Lowest Common Ancestor of a Binary Search Tree
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]




Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
*/
package main

import (
	"fmt"
	"github.com/frankerzeng/golang_demo/Common"
)

type TreeNode = Common.TreeNode

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	p := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	q := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	fmt.Println(lowestCommonAncestor(root, p, q))
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	rst := &TreeNode{}
	rst.Val = checkNode(root, p.Val, q.Val)
	return rst
}

// 搜索树，所以是有序的
func checkNode(root *TreeNode, p, q int) int {
	if root.Val == p || root.Val == q {
		return root.Val
	}
	if root.Left != nil {
		if root.Val > p && root.Val > q { // 分布在根节点左侧
			return checkNode(root.Left, p, q)
		}
	}
	if root.Right != nil {
		if root.Val < p && root.Val < q { // 分布在根节点右侧
			return checkNode(root.Right, p, q)
		}
	}

	if root.Right != nil && root.Left != nil { // 分布在根节点两侧
		if (root.Val < q && root.Val > p) || (root.Val > q && root.Val < p) {
			return root.Val
		}
	}

	return -1
}
