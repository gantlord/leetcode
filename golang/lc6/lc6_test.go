package lc6

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"test 2", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test 3", args{"A", 1}, "A"},
		{"test 4", args{"AB", 1}, "AB"},
		{"test 5", args{"AB", 2}, "AB"},
		{"test 6", args{"A", 2}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
