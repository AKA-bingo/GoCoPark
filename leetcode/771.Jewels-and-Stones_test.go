package leetcode

import (
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{J: "aA", S: "aAAbbbb"}, want: 3},
		{name: "Example 2", args: args{J: "z", S: "ZZ"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumJewelsInStones(tt.args.J, tt.args.S); got != tt.want {
				t.Errorf("NumJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
