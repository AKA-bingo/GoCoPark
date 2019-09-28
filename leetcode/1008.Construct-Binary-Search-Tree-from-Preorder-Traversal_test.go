package leetcode

import (
	"math"
	"reflect"
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestBstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "Example", args: args{preorder: []int{1}}, want: &TreeNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBstFromPreorderVol1(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "Example", args: args{preorder: []int{1}}, want: &TreeNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BstFromPreorderBak(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BstFromPreorderVol1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bstBuild(t *testing.T) {
	type args struct {
		preorder  []int
		index     *int
		rootValue int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "Example", args: args{preorder: []int{1}, index: new(int), rootValue: math.MaxInt32}, want: &TreeNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstBuild(tt.args.preorder, tt.args.index, tt.args.rootValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}
