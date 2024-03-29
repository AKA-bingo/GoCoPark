package leetcode

import (
	"testing"
)

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{[][]int{
			{3, 0, 8, 4},
			{2, 4, 5, 7},
			{9, 2, 6, 3},
			{0, 3, 1, 0},
		}}, want: 35},
		{name: "Example 2", args: args{[][]int{
			{13, 47, 91, 34, 20, 33, 39, 22, 80, 62},
			{73, 97, 88, 51, 38, 36, 52, 75, 25, 99},
			{95, 43, 32, 26, 82, 74, 60, 69, 59, 55},
			{20, 41, 77, 95, 79, 46, 70, 50, 17, 51},
			{51, 0, 93, 27, 46, 41, 58, 49, 8, 5},
			{92, 58, 38, 56, 73, 93, 34, 47, 23, 62},
			{97, 66, 57, 72, 26, 46, 4, 90, 82, 74},
			{7, 44, 67, 96, 0, 82, 75, 22, 53, 100},
			{95, 48, 46, 68, 41, 53, 69, 42, 13, 87},
			{79, 48, 96, 39, 21, 35, 3, 12, 22, 42},
		}}, want: 3673},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIncreaseKeepingSkyline(tt.args.grid); got != tt.want {
				t.Errorf("MaxIncreaseKeepingSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
