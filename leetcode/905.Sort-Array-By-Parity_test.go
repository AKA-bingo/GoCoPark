package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example", args: args{A: []int{3, 1, 2, 4}}, want: []int{4, 2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrayByParity(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
