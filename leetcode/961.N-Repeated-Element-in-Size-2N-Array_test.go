package leetcode

import (
	"testing"
)

func TestRepeatedNTimes(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{A: []int{1, 2, 3, 3}}, want: 3},
		{name: "Example 2", args: args{A: []int{2, 1, 2, 5, 3, 2}}, want: 2},
		{name: "Example 3", args: args{A: []int{5, 1, 5, 2, 5, 3, 5, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatedNTimes(tt.args.A); got != tt.want {
				t.Errorf("RepeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
