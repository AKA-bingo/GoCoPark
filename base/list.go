package base

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeInit(vals []int) *ListNode {
	root := new(ListNode)
	temp := root
	for index, val := range vals {
		temp.Val = val
		if index < len(vals)-1 {
			temp.Next = new(ListNode)
			temp = temp.Next
		}
	}
	return root
}

func (listNode *ListNode) Add(val int) {
	temp := listNode
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (listNode *ListNode) Print() string {
	res := ""
	temp := listNode
	for temp != nil {
		res += fmt.Sprintf("%v->", temp.Val)
		temp = temp.Next
	}
	res += fmt.Sprintf("NULL")
	return res
}
