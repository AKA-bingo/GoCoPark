package leetcode

import (
	"reflect"
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestInsertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	var tests []struct {
		name string
		args args
		want *TreeNode
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
