package leetcode

import (
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{[]int{0, 1, 0}}, want: 1},
		{name: "Example 1", args: args{[]int{0, 2, 1, 0}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PeakIndexInMountainArray(tt.args.A); got != tt.want {
				t.Errorf("PeakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
