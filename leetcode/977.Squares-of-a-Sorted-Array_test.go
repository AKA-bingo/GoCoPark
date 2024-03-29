package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example 1", args: args{[]int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{name: "Example 2", args: args{[]int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedSquares(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
