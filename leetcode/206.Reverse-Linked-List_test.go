package leetcode

import (
	"reflect"
	"testing"

	. "github.com/AKA-bingo/GoCoPark/base"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	var tests []struct {
		name string
		args args
		want *ListNode
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
