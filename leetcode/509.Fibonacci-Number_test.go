package leetcode

import (
	"testing"
)

func TestFib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{N: 2}, want: 1},
		{name: "Example 2", args: args{N: 3}, want: 2},
		{name: "Example 3", args: args{N: 4}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.N); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
