package leetcode

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "Example 1", args: args{A: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}}, want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{name: "Example 2", args: args{A: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}}, want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipAndInvertImage(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
