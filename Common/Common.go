package Common

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Person struct {
	name string
}

func (person *Person) Name() string {
	return person.name
}
func (person *Person) setName(name string) {
	person.name = name
}

func PrintList(head *ListNode) {
	currentNode := head
	var intSlice []int
	for currentNode != nil {
		intSlice = append(intSlice, currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println(intSlice)
}
