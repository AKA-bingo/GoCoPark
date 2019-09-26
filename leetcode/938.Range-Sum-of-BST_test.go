package leetcode

import (
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestRangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		L    int
		R    int
	}
	var tests []struct {
		name string
		args args
		want int
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeSumBST(tt.args.root, tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("RangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
