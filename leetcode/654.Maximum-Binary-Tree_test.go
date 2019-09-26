package leetcode

import (
	"reflect"
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests []struct {
		name string
		args args
		want *TreeNode
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
