package base

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeInit(val int) *ListNode {
	return &ListNode{Val: val}
}

func (this *ListNode) Add(node *ListNode) {
	temp := this
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
}

func (this *ListNode) Print() {
	temp := this
	for temp != nil {
		fmt.Printf("%v->", temp.Val)
		temp = temp.Next
	}
	fmt.Println("NULL")
}
