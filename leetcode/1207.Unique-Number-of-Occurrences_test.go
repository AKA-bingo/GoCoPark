package leetcode

import (
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example 1", args: args{[]int{1, 2, 2, 1, 1, 3}}, want: true},
		{name: "Example 2", args: args{[]int{1, 2}}, want: false},
		{name: "Example 3", args: args{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("UniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
