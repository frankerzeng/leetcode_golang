/*
655. Print Binary Tree
Print a binary tree in an m*n 2D string array following these rules:

1.The row number m should be equal to the height of the given binary tree.
2.The column number n should always be an odd number.
3.The root node's value (in string format) should be put in the exactly middle of the first row it can be put. The column
	and the row where the root node belongs will separate the rest space into two parts (left-bottom part and right-bottom part).
	You should print the left subtree in the left-bottom part and print the right subtree in the right-bottom part. The
	left-bottom part and the right-bottom part should have the same size. Even if one subtree is none while the other is not,
	you don't need to print anything for the none subtree but still need to leave the space as large as that for the other subtree.
	However, if two subtrees are none, then you don't need to leave space for both of them.
4.Each unused space should contain an empty string "".
5.Print the subtrees following the same rules.

Example 1:
Input:
     1
    /
   2
Output:
	[["", "1", ""],
	 ["2", "", ""]]

Example 2:
Input:
     1
    / \
   2   3
    \
     4
Output:
	[["", "", "", "1", "", "", ""],
	 ["", "2", "", "", "", "3", ""],
	 ["", "", "4", "", "", "", ""]]

Example 3:
Input:
      1
     / \
    2   5
   /
  3
 /
4
Output:
	[["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
	 ["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
	 ["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
	 ["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
Note: The height of binary tree is in the range of [1, 10].
 */
package main

import (
	"fmt"
	"../Common"
	"strconv"
)

type TreeNode = Common.TreeNode

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(printTree(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	var rst [][]string
	level := 0 // level of root
	var rootList = []*TreeNode{root}
	for len(rootList) > 0 {
		level++
		var tmp []*TreeNode
		for _, v := range rootList {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		rootList = tmp
	}

	// pow(2,level) length of level
	m := 1
	for i := 0; i < level; i++ {
		m *= 2
	}
	m -= 1

	for i := 0; i < level; i++ {
		var tmpList []string
		for j := 0; j < m; j++ {
			tmpList = append(tmpList, "")
		}
		rst = append(rst, tmpList)
	}

	change(&rst, root, 0, 0, m-1)
	return rst
}

// set middle of start - end to node.val
func change(rst *[][]string, node *TreeNode, currentLevel int, start int, end int) {
	keyIndex := (end + start) / 2
	(*rst)[currentLevel][keyIndex] = strconv.Itoa(node.Val)
	if node.Left != nil {
		change(rst, node.Left, currentLevel+1, start, keyIndex-1)
	}
	if node.Right != nil {
		change(rst, node.Right, currentLevel+1, keyIndex+1, end)
	}
}
