/*
99. Recover Binary Search Tree

Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Example 1:

Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2
Example 2:

Input: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
Follow up:

A solution using O(n) space is pretty straight forward.
Could you devise a constant space solution?z
*/
package main

import (
	"fmt"
	"github.com/frankerzeng/golang_demo/Common"
)

type TreeNode = Common.TreeNode

func main() {

	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(recoverTree(root))
}

func printNode(root *TreeNode) {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) *TreeNode {
	level := []*TreeNode{root}

	for true {
		for i := 0; i < len(level); i++ {
			rstTmp := leftValidate(level[i], level[i].Val)
			fmt.Println(rstTmp, "---rstTmp")
			if rstTmp.Val != -99 {
				firstVal := level[i].Val
				level[i].Val = rstTmp.Val
				println(level[i].Val, i)
				rstTmp.Val = firstVal
				fmt.Println(root, "---112")
				fmt.Println(root.Right, "---112")
				fmt.Println(root.Right.Right, "---112")

				return root
			}
			rstTmp = rightValidate(level[i], level[i].Val)

			fmt.Println(*rstTmp, "---2")
			if rstTmp.Val != -99 {
				firstVal := level[i].Val
				level[i].Val = rstTmp.Val
				println(level[i].Val, i)
				rstTmp.Val = firstVal

				fmt.Println(root, "---112")
				fmt.Println(root.Left, "---112")
				fmt.Println(root.Right, "---112")
				fmt.Println(root.Right.Right, "---112")

				return root
			}
		}
	}

	return root
}
func leftValidate(node *TreeNode, val int) *TreeNode { // val is bigger then node.val
	returnNode := &TreeNode{Val: -99}
	fmt.Println(returnNode, "returnNode")
	if node != nil {
		if node.Left != nil && node.Left.Val > val {
			returnNode = node.Left
		} else if node.Right != nil && node.Right.Val > val {
			returnNode = node.Right
		} else {
			if node.Left != nil {
				rst := leftValidate(node.Left, val)
				if rst.Val != -99 {
					return rst
				}
			}
			if node.Left != nil {
				rst := leftValidate(node.Right, val)
				if rst.Val != -99 {
					return rst
				}
			}
		}

	}
	return returnNode
}
func rightValidate(node *TreeNode, val int) *TreeNode { // val is bigger then node.val
	returnNode := &TreeNode{Val: -99}
	if node != nil {
		if node.Left != nil && node.Left.Val < val {
			returnNode = node.Left
		} else if node.Right != nil && node.Right.Val < val {
			returnNode = node.Right
		} else {
			if node.Left != nil {
				rst := rightValidate(node.Left, val)
				if rst.Val != -99 {
					return rst
				}
			}
			if node.Left != nil {
				rst := rightValidate(node.Right, val)
				if rst.Val != -99 {
					return rst
				}
			}
		}

	}
	return returnNode
}

// todo unfinish
