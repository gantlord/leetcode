package lc9

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{121}, true},
		{"t2", args{-121}, false},
		{"t3", args{10}, false},
		{"t4", args{-101}, false},
		{"t5", args{0}, true},
		{"t6", args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
