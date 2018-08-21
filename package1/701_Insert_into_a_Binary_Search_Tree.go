/*
701. Insert into a Binary Search Tree
Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST.
Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion.
You can return any of them.

For example,

	Given the tree:
			4
		   / \
		  2   7
		 / \
		1   3
	And the value to insert: 5
	You can return this binary search tree:

			 4
		   /   \
		  2     7
		 / \   /
		1   3 5
	This tree is also valid:

			 5
		   /   \
		  2     7
		 / \
		1   3
			 \
			  4
*/
package main

import "fmt"
import "../Common"

type TreeNode = Common.TreeNode

func main() {

	var node1 = new(TreeNode)
	var node11 = new(TreeNode)
	var node12 = new(TreeNode)
	var node111 = new(TreeNode)
	var node112 = new(TreeNode)

	node1.Val = 4
	node1.Right = node12
	node1.Left = node11
	node12.Val = 7
	node11.Val = 2
	node11.Right = node112
	node11.Left = node111
	node111.Val = 1
	node112.Val = 3

	fmt.Println(insertIntoBST(node1, 5))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	tmp := new(TreeNode) // defined a empty TreeNode
	if root == nil {
		root.Val = val
	} else {
		// val in right side or left side of root node
		if root.Val < val {
			if root.Right == nil {
				tmp.Val = val
				root.Right = tmp
				return root
			}
			root.Right = insertIntoBST(root.Right, val)
		} else {
			if root.Left == nil {
				tmp.Val = val
				root.Left = tmp
				return root
			}
			root.Left = insertIntoBST(root.Left, val)
		}
	}
	return root
}
