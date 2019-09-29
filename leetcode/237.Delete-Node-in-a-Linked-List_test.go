package leetcode

import (
	"reflect"
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestDeleteNode(t *testing.T) {
	type args struct {
		root *ListNode
		node *ListNode
	}
	ExampleList1 := ListNodeInit([]int{4, 5, 1, 9})
	ExampleList2 := ListNodeInit([]int{4, 5, 1, 9})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "Example 1", args: args{root: ExampleList1, node: ExampleList1.Next} /*5*/, want: ListNodeInit([]int{4, 1, 9})},
		{name: "Example 2", args: args{root: ExampleList2, node: ExampleList2.Next.Next} /*1*/, want: ListNodeInit([]int{4, 5, 9})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if DeleteNode(tt.args.node); !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("DeleteNode() = %v, want %v", tt.args.node, tt.want)
			}
		})
	}
}
