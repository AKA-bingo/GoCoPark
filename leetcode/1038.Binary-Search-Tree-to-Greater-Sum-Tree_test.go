package leetcode

import (
	"reflect"
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestBstToGst(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	var tests []struct {
		name string
		args args
		want *TreeNode
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BstToGst(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}
