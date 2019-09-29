package leetcode

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example", args: args{S: "ababcbacadefegdehijhklij"}, want: []int{9, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartitionLabels(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
