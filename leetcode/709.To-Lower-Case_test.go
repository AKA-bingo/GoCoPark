package leetcode

import (
	"testing"
)

func TestToLowerCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example 1", args: args{str: "Hello"}, want: "hello"},
		{name: "Example 2", args: args{str: "here"}, want: "here"},
		{name: "Example 3", args: args{str: "LOVELY"}, want: "lovely"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCase(tt.args.str); got != tt.want {
				t.Errorf("ToLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
